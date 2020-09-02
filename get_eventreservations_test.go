package resengo_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetEventReservations(t *testing.T) {
	req := client.NewGetEventReservationsRequest()
	req.QueryParams().GuestName = "info.turnova@corsendonkhotels.com"
	// req.QueryParams().PageSize = 20
	// req.QueryParams().Page = 1
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
