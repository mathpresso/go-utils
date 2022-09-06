package cloudcdn_test

import (
	"encoding/base64"
	"log"
	"net/url"
	"time"

	cloudcdnsign "github.com/mathpresso/go-utils/urlsign/cloudcdn"
)

func Example_simple() {
	keyName := "my-key"                 // key name you defined
	base64Key := "your-base64-key-here" // bas64 key from GCP
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Fatalf("failed to decode key: %s", err)
	}

	testURL, err := url.Parse("https://example.com/test-image.jpg?key=value")
	if err != nil {
		log.Fatalf("failed to parse url: %s", err)
	}

	signClient := cloudcdnsign.New(keyName, key)

	// sign url with 1 hour expiration
	signedURL, err := signClient.Sign(*testURL, time.Now().Add(1*time.Hour))
	if err != nil {
		log.Fatalf("failed to sign url: %s", err)
	}

	log.Printf("signed url: %s\n", signedURL)
}
