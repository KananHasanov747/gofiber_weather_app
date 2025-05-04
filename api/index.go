package handler

import (
	"net/http"

	"github.com/KananHasanov747/gofiber_weather_app/configs"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := configs.NewApp()
	return adaptor.FiberApp(app)
}
