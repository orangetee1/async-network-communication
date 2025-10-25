package service

import (
	"async_communication/internal/model"
	"log"
	"os"
	"resty.dev/v3"
)

func GetLocation(location string, pipe chan<- *model.Locations) {
	// TODO: un hardcode values
	c := resty.New()
	defer c.Close()

	locationKey := os.Getenv("GRAPH_HOPPER_API_KEY")

	res, err := c.R().
		SetQueryParam("q", location).
		SetQueryParam("key", locationKey).
		SetResult(&model.Locations{}).
		Get("https://graphhopper.com/api/1/geocode")

	if err != nil {
		log.Fatal(err)
	}

	pipe <- res.Result().(*model.Locations)
}
