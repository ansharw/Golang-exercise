package models

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
	StatusWeather string `json:"status"`
}

type Weather_ struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
