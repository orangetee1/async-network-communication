package service

import (
	"async_communication/internal/model"
	"log"
	"os"
	"resty.dev/v3"
)

type errorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func GetLocation(location string, pipe chan<- *model.Locations) {
	// TODO: un hardcode values
	locationKey := os.Getenv("GEOAPIFY_API_KEY")

	c := resty.New()
	defer c.Close()

	res, err := c.R().
		SetQueryParam("apiKey", locationKey).
		SetQueryParam("text", location).
		SetQueryParam("lang", "ru").
		SetQueryParam("format", "json").
		SetResult(&model.Locations{}).
		Get("https://api.geoapify.com/v1/geocode/search")

	if err != nil {
		log.Fatal(err)
	}

	pipe <- res.Result().(*model.Locations)
}
