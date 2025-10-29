package model

type Hit struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Postcode  string  `json:"postcode"`
	Type      string  `json:"result_type"`
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}

type Locations struct {
	Hits []Hit `json:"results"`
}
