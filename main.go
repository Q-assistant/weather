package main

import (
	"fmt"
	"github.com/q-assistant/sdk"
	"github.com/q-assistant/sdk/express"
	"github.com/q-assistant/sdk/logger"
	"github.com/q-assistant/sdk/update"
	"log"
	"time"
	"weather/client"
)

func main() {
	cnf := map[string]interface{}{
		"key":      "a",
		"secret":   "b",
		"location": "c",
	}

	skill, err := sdk.NewSkill("q", "weather", "0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	config, err := skill.WithConfig(cnf)
	if err != nil {
		log.Fatal(err)
	}

	weatherApi := client.NewOpenWeatherMap(config.String("key"), config.String("secret"), config.String("location"))

	skill.AddHandler("forecast", func(logger *logger.Logger, data *sdk.Data, express *express.Express) {
		layout := "2006-01-02T15:04:05-07:00"

		now, err := time.Parse(layout, time.Now().Format(layout))
		if err != nil {
			logger.Error(err)
			return
		}

		date, err := time.Parse(layout, data.Parameters.Fields["date-time"].GetStringValue())
		if err != nil {
			logger.Error(err)
			return
		}

		w, err := weatherApi.GetForecast()
		if err != nil {
			logger.Error(err)
			return
		}

		for _, w := range w {
			fmt.Println(w.DateTime, now)
			if w.DateTime.After(date) {
				express.Talk(fmt.Sprintf("%s, with a temperature of %.0f degrees", w.Description, w.Temp))
				break
			}
		}
	})

	skill.OnConfigUpdate(func(update *update.Update) {
		weatherApi = client.NewOpenWeatherMap(config.String("key"), config.String("secret"), config.String("location"))
	})

	skill.Run()
}
