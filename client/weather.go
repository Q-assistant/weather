package client

import "time"

type Weather struct {
	Type        string
	DateTime    time.Time
	Description string
	Temp        float64
	TempFeeling float64
	TempMin     float64
	TempMax     float64
	Humidity    float64
}
