package resengo_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetEventReservations(t *testing.T) {
	req := client.NewGetEventReservationsRequest()
	req.QueryParams().GuestName = "Turnova"
	req.QueryParams().PageSize = 50
	req.QueryParams().Page = 1
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
