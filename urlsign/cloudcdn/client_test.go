package cloudcdn_test

import (
	"errors"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/mathpresso/go-utils/urlsign"
	cloudcdnsign "github.com/mathpresso/go-utils/urlsign/cloudcdn"
)

func Test_Sign(t *testing.T) {
	testKey := []byte("test")
	testKeyName := "my-key"
	testExpireTime := time.Date(2030, 1, 2, 3, 40, 50, 0, time.UTC)

	testSignClient := cloudcdnsign.New(testKeyName, testKey)

	testCases := []struct {
		name        string
		testURL     url.URL
		expectURL   string
		expectError error
	}{
		{
			name:        "query violation 'Expires'",
			testURL:     mustParse(t, "https://example.com/test-image.jpg?Expires=1234567890"),
			expectError: urlsign.ErrNotAllowedQuery,
		},
		{
			name:        "query violation 'KeyName'",
			testURL:     mustParse(t, "https://example.com/test-image.jpg?KeyName=something"),
			expectError: urlsign.ErrNotAllowedQuery,
		},
		{
			name:        "query violation 'Signature'",
			testURL:     mustParse(t, "https://example.com/test-image.jpg?Signature="),
			expectError: urlsign.ErrNotAllowedQuery,
		},
		{
			name:        "multiple query violation",
			testURL:     mustParse(t, "https://example.com/test-image.jpg?Expires=1234567890&KeyName=something"),
			expectError: urlsign.ErrNotAllowedQuery,
		},
		{
			name:      "valid url",
			testURL:   mustParse(t, "https://example.com/test-image.jpg"),
			expectURL: "https://example.com/test-image.jpg?Expires=1893555650&KeyName=my-key&Signature=krZHzea7_qNZnHzbC6cOHTkpIQk=",
		},
		{
			name:      "valid url with query",
			testURL:   mustParse(t, "https://example.com/test-image.jpg?width=125"),
			expectURL: "https://example.com/test-image.jpg?width=125&Expires=1893555650&KeyName=my-key&Signature=lZy-ZRZTKwizOzEoARz3oLXgDF0=",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotURL, err := testSignClient.Sign(tc.testURL, testExpireTime)
			if tc.expectError == nil && err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if tc.expectError != nil && err == nil {
				t.Fatalf("expected error %s but got nil", tc.expectError)
			}
			if tc.expectError != nil && !errors.Is(err, tc.expectError) {
				t.Fatalf("expected error %s, but got %s", tc.expectError, err)
			}
			if strings.Compare(tc.expectURL, gotURL) != 0 {
				t.Errorf("expected url %s, but got %s", tc.expectURL, gotURL)
			}
		})
	}
}

func mustParse(t *testing.T, s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("failed to parse url: %s", err)
	}
	return *u
}
