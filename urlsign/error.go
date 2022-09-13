package urlsign

import "errors"

var (
	ErrNotAllowedQuery = errors.New("not allowed query parameter included")
)
