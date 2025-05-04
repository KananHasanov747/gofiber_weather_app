package handlers

import (
	"log"

	"github.com/KananHasanov747/gofiber_weather_app/app/services"
	"github.com/gofiber/fiber/v2"
)

// @Description Get list of cities by given City name.
// @Accept json
// @Produce json
// @Param city path string true "City Name"
// @Sucess 200 {object}
// @Router /api/v1/search/{city} [get]
func GetSearch(c *fiber.Ctx) error {
	city := c.Params("city")

	list, err := services.GetSearchData(city)
	if err != nil {
		log.Printf("GetSearch error: %v\n", err)
		return c.Status(500).SendString("GetSearch: upstream error")
	}

	c.Set("Content-Encoding", "gzip")
	c.Set("Content-Type", "application/json")
	return c.Send(list)
}

// @Description Get data of location by given Lat & Lon.
// @Accept json
// @Produce json
// @Param city path string true "City"
// @Param country path string true "Country"
// @Param lat path string true "Latitude"
// @Param lon path string true "Longitude"
// @Sucess 200 {object}
// @Router /api/v1/weather/{city}/{country}/{lat}/{lon} [get]
func GetWeather(c *fiber.Ctx) error {
	city, country, lat, lon := c.Params("city"), c.Params("country"), c.Params("lat"), c.Params("lon")

	data, err := services.GetWeatherData(city, country, lat, lon)
	if err != nil {
		log.Printf("GetWeather error: %v\n", err)
		return c.Status(500).SendString("GetWeather: upstream error")
	}

	c.Set("Content-Encoding", "gzip")
	c.Set("Content-Type", "application/json")
	return c.Send(data)
}
