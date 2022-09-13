package urlsign

import (
	"net/url"
	"time"
)

// Client is a helper for sign url with specific policy
type Client interface {
	// Sign signs the given url and returns signed-url as string.
	// We use string type because, signed-url is immutable.
	// given url.URL must not include query parameters which used for signing.
	Sign(originalURL url.URL, expiration time.Time) (string, error)
}
