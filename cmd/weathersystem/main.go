package main

import (
	"fmt"
	"net/http"

	"github.com/gabscristofani/weather-api/configs"
	"github.com/gabscristofani/weather-api/internal/adapters/api"
	"github.com/gabscristofani/weather-api/internal/infra/web"
	"github.com/gabscristofani/weather-api/internal/infra/web/webserver"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	locationClient := api.NewLocationClient(configs.LocationClientUrl)
	weatherClient := api.NewWeatherClient(configs.WeatherClientUrl, configs.WeatherClientKey)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webTempHandler := web.NewWebTempHandler(locationClient, weatherClient)
	webserver.AddHandler("/temp", http.MethodGet, webTempHandler.Get)
	fmt.Println("Starting web server on port ", configs.WebServerPort)
	webserver.Start()
}
