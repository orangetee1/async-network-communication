package model

type Hit struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}

type Locations struct {
	Hits []Hit `json:"results"`
}
