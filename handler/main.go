package handler

import (
	"fmt"
	"github.com/q-assistant/sdk"
	"github.com/q-assistant/sdk/express"
	"github.com/q-assistant/sdk/logger"
	"weather/client"
)

const timeLayout = "2006-01-02 15:04:05"

type Handler struct {
	client *client.OpenWeatherMap
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SetClient(client *client.OpenWeatherMap) {
	h.client = client
}

func (h *Handler) Forecast(logger *logger.Logger, data *sdk.Data, express *express.Express) {
	date, err := getDate(data.Parameters.Fields["datetimeV2"].GetListValue())
	if err != nil {
		logger.Error(err)
		return
	}

	w, err := h.client.GetForecast()
	if err != nil {
		logger.Error(err)
		return
	}

	for _, w := range w {
		if w.DateTime.After(date.Start) {
			express.Talk(fmt.Sprintf("%s, with a temperature of %.0f degrees", w.Description, w.Temp))
			break
		}
	}
}
