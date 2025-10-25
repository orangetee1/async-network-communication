package service

import (
	"async_communication/internal/model"
	"fmt"
	"log"
	"os"
	"resty.dev/v3"
)

func GetPlacesByRadius(lon float32, lat float32, radius int) chan *model.Places {
	pipe := make(chan *model.Places)

	go func() {
		placesKey := os.Getenv("GEOAPIFY_API_KEY")

		filter := "circle:" + fmt.Sprintf("%f", lon) +
			"," + fmt.Sprintf("%f", lat) +
			"," + fmt.Sprintf("%d", radius)

		c := resty.New()
		defer c.Close()

		res, err := c.R().
			SetQueryParam("apiKey", placesKey).
			SetQueryParam("categories", "entertainment").
			SetQueryParam("filter", filter).
			SetResult(&model.Places{}).
			Get("https://api.geoapify.com/v2/places")

		if err != nil {
			log.Fatal(err)
		}

		pipe <- res.Result().(*model.Places)

		close(pipe)
	}()

	return pipe
}
