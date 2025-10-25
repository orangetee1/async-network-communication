package model

type contact struct {
	Phone string `json:"phone"`
}

type properties struct {
	Name         string  `json:"name"`
	Website      string  `json:"website"`
	OpeningHours string  `json:"opening_hours"`
	Street       string  `json:"street"`
	City         string  `json:"city"`
	HouseNumber  string  `json:"housenumber"`
	Contact      contact `json:"contact"`
}

type feature struct {
	Properties properties `json:"properties"`
}

type PlaceInfo struct {
	Features []feature `json:"features"`
}
