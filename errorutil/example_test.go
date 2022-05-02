package errorutil_test

import (
	"errors"

	"github.com/mathpresso/go-utils/errorutil"
)

// If you want to just wrap error with stack trace, simply wrap your error with .Wrap()
func Example_simple() {
	_ = func() error {
		err := errors.New("some error")
		return errorutil.Wrap(err)
	}
}

// If you want to set some cause-error for your error, simply use .FromCause() option
func Example_nested() {
	_ = func() error {
		ErrStatic := errors.New("static error")
		causeErr := errors.New("cause error")
		return errorutil.Wrap(ErrStatic, errorutil.FromCause(causeErr))
	}
}
