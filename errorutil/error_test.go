package errorutil

import (
	"errors"
	"reflect"
	"testing"
)

func TestErrorMethod(t *testing.T) {
	rawErr := errors.New("test error")
	wrappedErr := Wrap(rawErr)
	if rawErr.Error() != wrappedErr.Error() {
		t.Errorf(".Error() mismatch:\nexpect=%#v\ngot=%#v", rawErr, wrappedErr)
	}
}

func TestErrorCompare(t *testing.T) {
	rawRoot := errors.New("root error")
	rawChild := errors.New("child error")

	// wrapped error must be equal to original error
	wrappedChild := Wrap(rawChild)
	if !errors.Is(wrappedChild, rawChild) {
		t.Error("wrappedChild is not rawChild")
	}

	// wrapped error with cause must be equal to both original error and original cause error
	wrappedRoot := Wrap(rawRoot, FromCause(rawChild))
	if !errors.Is(wrappedRoot, rawRoot) {
		t.Error("wrappedRoot is not rawRoot")
	}
	if !errors.Is(wrappedRoot, rawChild) {
		t.Error("wrappedRoot is not rawChild")
	}

	// double-wrapped error also
	doubleWrappedRoot := Wrap(rawRoot, FromCause(wrappedChild))
	if !errors.Is(doubleWrappedRoot, rawRoot) {
		t.Error("doubleWrappedRoot is not rawRoot")
	}
	if !errors.Is(doubleWrappedRoot, rawChild) {
		t.Error("doubleWrappedRoot is not rawChild")
	}
}

func TestStackTraceBind(t *testing.T) {
	// This test simulates some steps of extracting stack trace from https://github.com/getsentry/sentry-go
	// stacktrace_test.go will ensure traceable has proper stack trace.
	err := Wrap(errors.New("test"), AutoStackTrace())

	// check has StackTrace method
	methodStackTrace := reflect.ValueOf(err).MethodByName("StackTrace")
	if !methodStackTrace.IsValid() {
		t.Error("err.StackTrace() is invalid")
		return
	}

	// ensure StackTrace method returns slice type
	tc := methodStackTrace.Call(make([]reflect.Value, 0))[0]
	if tc.Kind() != reflect.Slice {
		t.Error("err.StackTrace() must return slice type")
		return
	}

	// ensure all items of slice is uintptr type
	for i := 0; i < tc.Len(); i++ {
		pc := tc.Index(i)
		if pc.Kind() != reflect.Uintptr {
			t.Error("all items of err.StackTrace() must be uintptr")
			return
		}
	}
}
