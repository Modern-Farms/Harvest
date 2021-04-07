package models

type GerminatorSettings struct {
	IsAuto       float64 `json:"is_auto"`
	HumidityLow  float64 `json:"humidity_low"`
	HumidityHigh float64 `json:"humidity_high"`
	LightOnTime  float64 `json:"light_on_time"`
	LightOffTime float64 `json:"light_off_time"`
	Mister       int     `json:"mister"`
	Fan          int     `json:"fan"`
	LED          int     `json:"led"`
}
