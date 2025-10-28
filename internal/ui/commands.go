package ui

import (
	"async_communication/internal/model"
	"async_communication/internal/service"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func RequestLocations(request string) tea.Cmd {
	return func() tea.Msg {
		locations := service.GetLocations(request)

		if locations.IsError() {
			err := locations.Error().(*model.LocationError)

			return ErrorChanged{
				Error: fmt.Sprintf("Code: %d, Message: %s", err.StatusCode, err.Message),
			}
		}

		return LocationsLoaded{
			locations: locations.Result().(*model.Locations),
		}
	}
}

func RequestWeather(latitude, longitude float32) tea.Cmd {
	return func() tea.Msg {
		weather := service.GetWeather(latitude, longitude)

		if weather.IsError() {
			err := weather.Error().(*model.WeatherError)

			return ErrorChanged{
				Error: fmt.Sprintf("Code %d, Message: %s", err.StatusCode, err.Message),
			}
		}

		return WeatherLoaded{
			weather: weather.Result().(*model.Weather),
		}
	}
}
