package errorutil

import (
	"fmt"
	"strings"
	"testing"
)

func simpleTraceable(s int) *traceable {
	return traceableFromCallers(s)
}

func traceableFuncBuilder(s int) func() *traceable {
	return func() *traceable {
		return traceableFromCallers(s)
	}
}

func traceableCallWrapper(s int) *traceable {
	return traceableFuncBuilder(s)()
}

func checkStackTrace(t *testing.T, given *traceable, expect []string) {
	traceString := fmt.Sprintf("%+v", given.StackTrace())
	traceStringParts := make([]string, 0, len(expect))
	for idx, traceLine := range strings.SplitN(strings.Trim(traceString, " \n\r\t"), "\n", len(expect)*2+1) {
		if idx%2 != 0 {
			continue
		}
		if idx >= len(expect)*2 {
			break
		}
		traceStringParts = append(traceStringParts, traceLine)
	}
	for i := 0; i < len(expect); i++ {
		if !strings.HasPrefix(traceStringParts[i], expect[i]) {
			t.Errorf("stacktrace mismatch:\nexpect=%#v\ngot=%#v", expect, traceStringParts)
			break
		}
	}
}

func TestProperStackTrace(t *testing.T) {
	testCases := []struct {
		name     string
		runnable func() *traceable
		pattern  []string
	}{
		{
			name: "SimpleTrace",
			runnable: func() *traceable {
				return simpleTraceable(1)
			},
			pattern: []string{
				"github.com/mathpresso/go-utils/errorutil.traceableFromCallers",
				"github.com/mathpresso/go-utils/errorutil.simpleTraceable",
				"github.com/mathpresso/go-utils/errorutil.TestProperStackTrace.",
			},
		},
		{
			name: "NestedTrace",
			runnable: func() *traceable {
				return traceableCallWrapper(2)
			},
			pattern: []string{
				"github.com/mathpresso/go-utils/errorutil.traceableFuncBuilder.",
				"github.com/mathpresso/go-utils/errorutil.traceableCallWrapper",
				"github.com/mathpresso/go-utils/errorutil.TestProperStackTrace.",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			checkStackTrace(t, tt.runnable(), tt.pattern)
		})
	}
}
