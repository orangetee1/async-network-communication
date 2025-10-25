package model

type Hit struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Name    string `json:"name"`
}

type Locations struct {
	Hits []Hit `json:"hits"`
}
