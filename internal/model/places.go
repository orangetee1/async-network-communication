package model

type placeProperties struct {
	Id        string `json:"place_id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Formatted string `json:"formatted"`
}

type place struct {
	Properties placeProperties `json:"properties"`
}

type Places struct {
	Features []place `json:"features"`
}
