package main

import (
	"github.com/q-assistant/sdk"
	"github.com/q-assistant/sdk/config"
	"github.com/q-assistant/sdk/update"
	"log"
	"weather/client"
	"weather/handler"
)

func main() {
	cnf := map[string]interface{}{
		"key":    "a",
		"secret": "a",
		"location": map[string]interface{}{
			"lat":  0,
			"lon":  0,
			"name": "a",
		},
	}

	skill, err := sdk.NewSkill("q", "weather", "0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	config, err := skill.WithConfig(cnf)
	if err != nil {
		log.Fatal(err)
	}

	handler := handler.New()
	handler.SetClient(createWeatherApi(config))


	skill.AddHandler("forecast", handler.Forecast)

	skill.OnConfigUpdate(func(update *update.Update) {
		handler.SetClient(createWeatherApi(config))
	})

	skill.Run()
}

func createWeatherApi(config config.Config) *client.OpenWeatherMap {
	location := config.Map("location")
	return client.NewOpenWeatherMap(&client.Config{
		Key:    config.String("key"),
		Secret: config.String("secret"),
		Location: &client.Location{
			Lat:  location["lat"].(float64),
			Lon:  location["lon"].(float64),
			Name: location["name"].(string),
		},
	})

}