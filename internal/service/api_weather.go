package service

import (
	"async_communication/internal/model"
	"fmt"
	"log"
	"os"
	"resty.dev/v3"
)

func GetWeather(latitude, longitude float32) *resty.Response {
	c := resty.New()
	defer c.Close()

	weatherKey := os.Getenv("OPEN_WEATHER_API_KEY")

	res, err := c.R().
		SetQueryParam("lat", fmt.Sprintf("%f", latitude)).
		SetQueryParam("lon", fmt.Sprintf("%f", longitude)).
		SetQueryParam("units", "metric").
		SetQueryParam("appid", weatherKey).
		SetResult(&model.Weather{}).
		SetError(&model.WeatherError{}).
		Get("https://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		log.Fatal(err)
	}

	return res
}
