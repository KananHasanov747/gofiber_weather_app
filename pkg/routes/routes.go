package routes

import (
	"github.com/KananHasanov747/gofiber_weather_app/app/handlers"
	"github.com/KananHasanov747/gofiber_weather_app/app/services"
	"github.com/KananHasanov747/gofiber_weather_app/pkg/utils"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/swagger")

	route.Get("*", swagger.HandlerDefault)
}

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/search/:city", handlers.GetSearch).Name("search")                       // get a list of cities
	route.Get("/weather/:city/:country/:lat/:lon", handlers.GetWeather).Name("weather") // get a data from coordinates
}

func TemplateRoutes(route *fiber.App) {
	route.Get("/", func(c *fiber.Ctx) error {
		var city, country, lat, lon string
		loc, err := utils.GetLocation(c.IP())

		if err != nil {
			city, country, lat, lon = "Baku", "Azerbaijan", "40.37767", "49.89201"
		} else {
			city, country, lat, lon = loc["city"], loc["country"], loc["latitude"], loc["longitude"]
		}

		raw, _ := services.GetWeatherData(city, country, lat, lon)
		data, _ := utils.GzipDecompress(raw)

		return c.Render("index", fiber.Map{"data": data})
	})
}
