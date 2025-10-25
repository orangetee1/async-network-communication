package model

type WeatherDesc struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type WeatherValues struct {
	Temp     float32 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

type WindInfo struct {
	Speed float32 `json:"speed"`
}

type Weather struct {
	Description []WeatherDesc `json:"weather"`
	Values      WeatherValues `json:"main"`
	Wind        WindInfo      `json:"wind"`
}
