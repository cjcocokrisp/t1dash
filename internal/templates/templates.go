package templates

import (
	"embed"
	"html/template"

	log "github.com/sirupsen/logrus"
)

//go:embed web/templates/*.html
var templateFS embed.FS

var Templates *template.Template

// InitTemplates inits a exported variable that holds all
// webpage templates. This function must be ran before
// any of the handlers are called
func InitTemplates() {
	t, err := template.ParseFS(templateFS, "web/templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
	Templates = t
}
