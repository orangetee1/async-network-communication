package service

import (
	"async_communication/internal/model"
	"fmt"
	"log"
	"os"
	"resty.dev/v3"
)

func GetPlacesByRadius(lon float32, lat float32, radius int) *resty.Response {
	placesKey := os.Getenv("GEOAPIFY_API_KEY")

	filter := fmt.Sprintf("circle:%f,%f,%d", lon, lat, radius)

	c := resty.New()
	defer c.Close()

	res, err := c.R().
		SetQueryParam("apiKey", placesKey).
		SetQueryParam("categories", "entertainment").
		SetQueryParam("filter", filter).
		SetResult(&model.Places{}).
		SetError(&model.LocationError{}).
		Get("https://api.geoapify.com/v2/places")

	if err != nil {
		log.Fatal(err)
	}

	return res
}
