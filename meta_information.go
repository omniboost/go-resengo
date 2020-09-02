package resengo

type MetaInformation struct {
	TotalResources int `json:"@TotalResources"`
	TotalPages     int `json:"@TotalPages"`
	CurrentPage    int `json:"@CurrentPage"`
}
