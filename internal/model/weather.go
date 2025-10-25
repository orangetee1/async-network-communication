package model

type weatherDesc struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type weatherValues struct {
	Temp     float32 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

type windInfo struct {
	Speed float32 `json:"speed"`
}

type Weather struct {
	Description []weatherDesc `json:"weather"`
	Values      weatherValues `json:"main"`
	Wind        windInfo      `json:"wind"`
}
