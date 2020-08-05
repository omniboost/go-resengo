package fortnox_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	fortnox "github.com/omniboost/go-fortnox"
)

var (
	client    *fortnox.Client
	companyID int
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	clientSecret := os.Getenv("CLIENT_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	debug := os.Getenv("DEBUG")

	client = fortnox.NewClient(nil, clientSecret, accessToken)
	if debug != "" {
		client.SetDebug(true)
	}
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	client.SetDisallowUnknownFields(true)
	m.Run()
}
