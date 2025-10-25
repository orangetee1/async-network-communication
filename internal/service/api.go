package service

import (
	"log"
	"os"
	"resty.dev/v3"
)

type Hit struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Name    string `json:"name"`
}
type Locations struct {
	Hits []Hit `json:"hits"`
}

func GetLocation(location string, pipe chan<- *Locations) {
	// TODO: un hardcode values
	c := resty.New()
	defer c.Close()

	locationKey := os.Getenv("GRAPH_HOPPER_API_KEY")

	res, err := c.R().
		SetQueryParam("q", location).
		SetQueryParam("key", locationKey).
		SetResult(&Locations{}).
		Get("https://graphhopper.com/api/1/geocode")

	if err != nil {
		log.Fatal(err)
	}

	pipe <- res.Result().(*Locations)
}
