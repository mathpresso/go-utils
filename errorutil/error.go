package errorutil

import (
	"errors"
	"fmt"
	"io"
)

var _ error = (*wrapped)(nil)
var _ fmt.Formatter = (*wrapped)(nil)

type wrapped struct {
	error
	*causer
	*traceable
}

func (w *wrapped) Is(err error) bool {
	return errors.Is(w.error, err)
}

func (w *wrapped) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v':
		if f.Flag('+') {
			_, _ = fmt.Fprintf(f, "%s (caused by: %v)", w.Error(), w.Cause())
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(f, w.Error())
	}
}

type wrapOpt func(w *wrapped)

// AutoStackTrace bind stack trace from Wrap method caller to the error.
// You don't have to provide this option, because this is used by default.
func AutoStackTrace() wrapOpt {
	return func(w *wrapped) {
		w.traceable = traceableFromCallers(4)
	}
}

// FromCause wrap the error with provided cause. If you Unwrap this error, provided cause will be extracted.
func FromCause(err error) wrapOpt {
	return func(w *wrapped) {
		w.causer = &causer{cause: err}
	}
}

// Wrap wraps the error with provided opts.
func Wrap(err error, opts ...wrapOpt) error {
	if err == nil {
		return nil
	}
	if we, ok := err.(*wrapped); ok && len(opts) == 0 {
		// If error is already wrapped, and no additional options provided, just return it
		return we
	}

	w := &wrapped{error: err}
	for _, opt := range opts {
		opt(w)
	}
	if w.traceable == nil {
		// Auto bind stack trace if not already set
		AutoStackTrace()(w)
	}

	return w
}
