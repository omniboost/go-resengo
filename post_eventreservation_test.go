package resengo_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	resengo "github.com/omniboost/go-resengo"
)

func TestPostEventReservations(t *testing.T) {
	req := client.NewPostEventReservationsRequest()

	startTime := time.Now()
	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 19, 30, 0, 0, time.Local)
	endTime := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour()+2, startTime.Minute(), startTime.Second(), startTime.Nanosecond(), time.Local)
	body := resengo.PostEventReservationsRequestBody{
		TimeZone:             "Europe/Brussels",
		StartTimeLocal:       resengo.LocalTime{startTime},
		EndTimeLocal:         resengo.LocalTime{endTime},
		EventCategoryID:      1317,
		PartnerID:            51,
		NumberOfPersons:      2,
		CommunicationStatus:  "Delivered",
		EmailReceptionStatus: "Delivered",
		SMSReceptionStatus:   "NoSMS",
		GuestSummary: resengo.GuestSummary{
			PersonID: 13551855,
		},
		GuestRemarks: resengo.GuestRemarks{{
			PersonID: 13551855,
			Owner:    "Guest",
			Remark:   "Gastronomisch arrangement Turnova\r\nNaam van de gast: Bilaey\r\ninbegrepen: 5 gangenmenu\r\nEnkel het 5 gangenmenu op factuur naar Turnova.\r\nDranken en andere ter plaatse door de gast zelf.",
		}},
	}
	req.SetRequestBody(body)
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
