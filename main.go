package main

import (
	"log"

	"github.com/KananHasanov747/gofiber_weather_app/configs"
	_ "github.com/KananHasanov747/gofiber_weather_app/docs"
)

func main() {
	app := configs.NewApp()

	log.Fatal(app.Listen(configs.BASE_URL))
}
