package model

type Hit struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
}

type Locations struct {
	Hits []Hit `json:"results"`
}
