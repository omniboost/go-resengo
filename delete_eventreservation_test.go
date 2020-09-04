package resengo_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDeleteEventReservations(t *testing.T) {
	req := client.NewDeleteEventReservationsRequest()
	req.PathParams().EventReservationID = 43160542
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
