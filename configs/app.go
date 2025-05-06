package configs

import (
	"github.com/KananHasanov747/gofiber_weather_app/cache/redis"
	"github.com/KananHasanov747/gofiber_weather_app/pkg/routes"
	"github.com/KananHasanov747/gofiber_weather_app/pkg/utils"
	"github.com/KananHasanov747/gofiber_weather_app/templates"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	BASE_URL        = utils.Show("0.0.0.0:8000", ":8000")
	DEBUG           = utils.Show("DEBUG", "false")
	ALLOWED_ORIGINS = utils.Show("ALLOWED_ORIGINS", "*")
)

func NewApp() *fiber.App {
	engine := templates.TemplateEngine()

	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: ALLOWED_ORIGINS,
	}))

	app.Static("/static", "./static")

	redis.Init()

	routes.TemplateRoutes(app)
	routes.PublicRoutes(app)
	routes.SwaggerRoute(app)

	return app
}
