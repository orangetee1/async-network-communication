package ui

import "async_communication/internal/model"

type (
	SwitchToDetails struct {
		Hit model.Hit
	}
	SwitchToSurvey struct{}
)

type ErrorChanged struct {
	Error string
}

type LocationsLoaded struct {
	Locations *model.Locations
}

type WeatherLoaded struct {
	Weather *model.Weather
}

type PlacesLoaded struct {
	Places *model.Places
}

type PlaceInfoLoaded struct {
	Info *model.PlaceInfo
}
