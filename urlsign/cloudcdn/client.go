package cloudcdn

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/mathpresso/go-utils/urlsign"
)

const (
	expireQueryKey    = "Expires"
	keyNameQueryKey   = "KeyName"
	signatureQueryKey = "Signature"
)

var _ urlsign.Client = (*client)(nil)

type client struct {
	keyName string
	key     []byte
}

func (c *client) Sign(originalURL url.URL, expiration time.Time) (string, error) {
	// check given url has query parameters which used for signing
	originalQuery := originalURL.Query()
	for _, k := range []string{expireQueryKey, keyNameQueryKey, signatureQueryKey} {
		if _, exist := originalQuery[k]; exist {
			return "", urlsign.ErrNotAllowedQuery
		}
	}

	signedURL := originalURL
	signQuery := url.Values{
		expireQueryKey:  []string{strconv.Itoa(int(expiration.Unix()))},
		keyNameQueryKey: []string{c.keyName},
	}

	// Add sign query
	if signedURL.RawQuery == "" {
		signedURL.RawQuery = signQuery.Encode()
	} else {
		signedURL.RawQuery = fmt.Sprintf("%s&%s", signedURL.RawQuery, signQuery.Encode())
	}

	// Generate signature of url
	mac := hmac.New(sha1.New, c.key)
	mac.Write([]byte(signedURL.String()))
	signature := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	signedURL.RawQuery = fmt.Sprintf("%s&%s=%s", signedURL.RawQuery, signatureQueryKey, signature)

	return signedURL.String(), nil
}

// New creates urlsign.Client for GCP Cloud CDN.
func New(keyName string, key []byte) urlsign.Client {
	return &client{
		keyName: keyName,
		key:     key,
	}
}
