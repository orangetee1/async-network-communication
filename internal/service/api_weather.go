package service

import (
	"async_communication/internal/model"
	"log"
	"os"
	"resty.dev/v3"
)

func GetWeather(latitude string, longitude string) chan *model.Weather {
	pipe := make(chan *model.Weather)

	go func() {
		c := resty.New()
		defer c.Close()

		weatherKey := os.Getenv("OPEN_WEATHER_API_KEY")

		res, err := c.R().
			SetQueryParam("lat", latitude).
			SetQueryParam("lon", longitude).
			SetQueryParam("units", "metric").
			SetQueryParam("appid", weatherKey).
			SetResult(&model.Weather{}).
			Get("https://api.openweathermap.org/data/2.5/weather")

		if err != nil {
			log.Fatal(err)
		}

		pipe <- res.Result().(*model.Weather)

		close(pipe)
	}()

	return pipe
}
