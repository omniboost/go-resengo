package fortnox

import "github.com/omniboost/go-fortnox/utils"

type Accounts []Account

type Account struct {
	URL                   utils.URL `json:"@url"`
	Active                bool      `json:"active"`
	BalanceBroughtForward float64   `json:"BalanceBroughtForward"`
	CostCenter            string    `json:"CostCenter"`
	CostCentersettings    string    `json:"CostCentersettings"`
	Description           string    `json:"Description"`
	Number                int       `json:"Number"`
	Project               string    `json:"Project"`
	ProjectSettings       string    `json:"ProjectSettings"`
	SRU                   int       `json:"SRU"`
	Year                  int       `json:"Year"`
	VATCode               string    `json:"VATCode"`
}
