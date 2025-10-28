package model

type LocationError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

type WeatherError struct {
	StatusCode int    `json:"cod"`
	Message    string `json:"message"`
}
