package errorutil

// causer provides error chaining mechanism with hierarchy
type causer struct {
	cause error
}

// Cause returns chained error
func (c *causer) Cause() error {
	if c == nil {
		return nil
	}
	return c.cause
}

// Unwrap stands for error chaining compatibility
func (c *causer) Unwrap() error {
	return c.Cause()
}
