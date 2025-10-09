package templates

import (
	"embed"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/cjcocokrisp/t1dash/pkg/util"
)

//go:embed web/templates/** web/static/*
var WebFS embed.FS

// Var that holds templates
var Templates *template.Template

// InitTemplates inits a exported variable that holds all
// webpage templates. This function must be ran before
// any of the handlers are called
func InitTemplates() {
	var err error
	Templates, err = parseEmbedTemplates()
	if err != nil {
		util.LogError("FATAL", "templates", err)
		os.Exit(1)
	}
}

// parseEmbedTemplates parses the WebFS embed filesystem for html
// templates and adds them to a template which is returned
func parseEmbedTemplates() (*template.Template, error) {
	t := template.New("")
	templateFS, err := fs.Sub(WebFS, "web/templates")
	if err != nil {
		return nil, err
	}

	err = fs.WalkDir(templateFS, ".", func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if dir.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".html" {
			content, err := fs.ReadFile(templateFS, path)
			if err != nil {
				return err
			}
			t, err = t.New(path).Parse(string(content))
		}
		return nil
	})
	return t, err
}
