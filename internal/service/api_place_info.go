package service

import (
	"async_communication/internal/model"
	"log"
	"os"
	"resty.dev/v3"
)

func GetPlaceInfoById(placeId string) chan *model.PlaceInfo {
	pipe := make(chan *model.PlaceInfo)

	go func() {
		placesKey := os.Getenv("GEOAPIFY_API_KEY")

		c := resty.New()
		defer c.Close()

		res, err := c.R().
			SetQueryParam("apiKey", placesKey).
			SetQueryParam("id", placeId).
			SetQueryParam("features", "details,details.names").
			SetResult(&model.PlaceInfo{}).
			Get("https://api.geoapify.com/v2/place-details")

		if err != nil {
			log.Fatal(err)
		}

		pipe <- res.Result().(*model.PlaceInfo)

		close(pipe)
	}()

	return pipe
}
