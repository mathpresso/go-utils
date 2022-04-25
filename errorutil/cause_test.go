package errorutil

import (
	"errors"
	"testing"
)

var _ error = (*errorWithCause)(nil)

// errorWithCause is simple custom error which includes causer
type errorWithCause struct {
	error
	*causer
}

func (e errorWithCause) Error() string {
	return e.error.Error()
}

func (e errorWithCause) Is(err error) bool {
	return errors.Is(e.error, err)
}

func TestNestedError(t *testing.T) {
	errRoot := errors.New("this is root error")
	errChild := errors.New("this is child error")
	errGrandChild := errors.New("this is grand child error")

	childErr := errorWithCause{
		error:  errChild,
		causer: &causer{cause: errGrandChild},
	}
	rootErr := errorWithCause{
		error:  errRoot,
		causer: &causer{cause: childErr},
	}

	// Ensure child has grandchild
	if valid := errors.Is(childErr, errGrandChild); !valid {
		t.Error("child don't have grandchild")
	}
	// Ensure root has child
	if valid := errors.Is(rootErr, errChild); !valid {
		t.Error("root don't have child")
	}
	// Ensure root has grandchild
	if valid := errors.Is(rootErr, errGrandChild); !valid {
		t.Error("root don't have grandchild")
	}
}
