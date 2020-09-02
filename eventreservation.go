package resengo

import "time"

type EventReservations []EventReservation

type EventReservation struct {
	ID                          int           `json:"id"`
	CompanyID                   int           `json:"companyId"`
	TimeZone                    string        `json:"timeZone"`
	StartTimeLocal              LocalTime     `json:"startTimeLocal"`
	EventCategoryID             int           `json:"eventCategoryId"`
	PartnerID                   int           `json:"partnerId"`
	NumberOfPersons             int           `json:"numberOfPersons"`
	StatusID                    int           `json:"statusId"`
	Index                       int           `json:"index"`
	EventID                     int           `json:"eventId"`
	ResourceIdsAreFixedChoice   bool          `json:"resourceIdsAreFixedChoice"`
	WaitingListPosition         int           `json:"waitingListPosition"`
	EndTimeLocal                LocalTime     `json:"endTimeLocal"`
	InternalRemark              string        `json:"internalRemark"`
	KitchenRemark               string        `json:"kitchenRemark"`
	Passage                     bool          `json:"passage"`
	LastGuestRemarkTreatDateUtc time.Time     `json:"lastGuestRemarkTreatDateUtc"`
	ExternalReservationID       string        `json:"externalReservationId"`
	CommunicationStatus         string        `json:"communicationStatus"`
	EmailReceptionStatus        string        `json:"emailReceptionStatus"`
	SmsReceptionStatus          string        `json:"smsReceptionStatus"`
	InsertDateUtc               time.Time     `json:"insertDateUtc"`
	UpdateDateUtc               time.Time     `json:"updateDateUtc"`
	GuestSummary                GuestSummary  `json:"guestSummary"`
	BookerSummary               BookerSummary `json:"bookerSummary"`
	GuestRemarks                GuestRemarks  `json:"guestRemarks"`
	Includes                    struct {
		Resources Resources `json:"resources"`
	} `json:"includes"`
	ManualAssignedResources bool `json:"manualAssignedResources"`
}

type GuestSummary struct {
	PersonID    int    `json:"personId"`
	GuestName   string `json:"guestName"`
	PhoneNumber string `json:"phoneNumber"`
	NoVisits    int    `json:"noVisits"`
	NoNoShows   int    `json:"noNoShows"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}

type BookerSummary struct {
	PersonID    int    `json:"personId"`
	BookerName  string `json:"bookerName"`
	PhoneNumber string `json:"phoneNumber"`
	NoVisits    int    `json:"noVisits"`
	NoNoShows   int    `json:"noNoShows"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}

type GuestRemarks []GuestRemark

type GuestRemark struct {
	ID            int       `json:"id"`
	ReservationID int       `json:"reservationId"`
	PersonID      int       `json:"personId"`
	Owner         string    `json:"owner"`
	UpdateDateUtc time.Time `json:"updateDateUtc"`
	Remark        string    `json:"remark"`
}

type Resources []Resource

type Resource struct {
	ResourceID      int    `json:"resourceId"`
	ResourceName    string `json:"resourceName"`
	NumberOfPersons int    `json:"numberOfPersons"`
	Remark          string `json:"remark"`
}
