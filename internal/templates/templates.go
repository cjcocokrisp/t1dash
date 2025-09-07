package templates

import (
	"embed"
	"html/template"
	"log"
)

//go:embed web/templates/*.html
var templateFS embed.FS

var Templates *template.Template

func InitTemplates() {
	t, err := template.ParseFS(templateFS, "web/templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates, %v", err)
	}
	Templates = t
}
