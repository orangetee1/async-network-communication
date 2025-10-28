package service

import (
	"async_communication/internal/model"
	"log"
	"os"
	"resty.dev/v3"
)

func GetPlaceInfoById(placeId string) *model.PlaceInfo {
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

	return res.Result().(*model.PlaceInfo)
}
