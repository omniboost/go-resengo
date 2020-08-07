package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	req := client.NewGetCustomerRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
