package middleware

import (
	"path/filepath"
	"strings"

	"github.com/KananHasanov747/gofiber_weather_app/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func ServeCompressedStatic(c *fiber.Ctx) error {
	filePath := c.Path() // e.g., /static/css/styles.css

	if strings.Contains(c.Get("Accept-Encoding"), "br") {
		if utils.PathExists("." + filePath + ".br") {
			c.Set("Content-Encoding", "br")
			c.Type(filepath.Ext(filePath))
			return c.SendFile("."+filePath+".br", false)
		}
	}
	if strings.Contains(c.Get("Accept-Encoding"), "gzip") {
		if utils.PathExists("." + filePath + ".gz") {
			c.Set("Content-Encoding", "gzip")
			c.Type(filepath.Ext(filePath))
			return c.SendFile("."+filePath+".gz", false)
		}
	}

	return c.Next()
}

// func RestrictAccessMiddleware(c *fiber.Ctx) error {
// 	RESTRICTED_URLS := []string{"/swagger"}
//
//
// 	return c.Next()
// }
