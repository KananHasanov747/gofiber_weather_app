package templates

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/gofiber/template/html/v2"
)

//go:embed *.html components/*.html
var templatesFS embed.FS

func TemplateEngine() *html.Engine {
	// development (checks if templates directory exists)
	if stat, err := os.Stat("templates"); err == nil && stat.IsDir() {
		return html.New("templates", ".html")
	}

	// production/serverless
	subFS, err := fs.Sub(templatesFS, ".")
	if err != nil {
		panic("failed to create sub FS for templates: " + err.Error())
	}

	return html.NewFileSystem(http.FS(subFS), ".html")
}
