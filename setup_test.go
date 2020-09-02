package resengo_test

import (
	"log"
	"os"
	"strconv"
	"testing"

	resengo "github.com/omniboost/go-resengo"
)

var (
	client    *resengo.Client
	companyID int
)

func TestMain(m *testing.M) {
	companyID, err := strconv.Atoi(os.Getenv("COMPANY_ID"))
	if err != nil {
		log.Fatal(err)
	}
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	debug := os.Getenv("DEBUG")

	client = resengo.NewClient(nil, companyID, clientID, clientSecret)
	if debug != "" {
		client.SetDebug(true)
	}
	m.Run()
}
