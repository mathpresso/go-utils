package errorutil

type causer struct {
	cause error
}

func (c *causer) Cause() error {
	if c == nil || c.cause == nil {
		return nil
	}
	return c.cause
}

// Unwrap stands for error chaining compatibility
func (c *causer) Unwrap() error {
	return c.Cause()
}
