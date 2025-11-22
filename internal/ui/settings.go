package ui

import (
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

// AppSettingsContent loads the content for app settings
func AppSettingsContent(w http.ResponseWriter, r *http.Request) {
	util.LogGetRequest("/settings/app", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "settings/app.html", nil)
	}
}

// GlucoseSettingsContent loads the content for general settings
func GlucoseSettingsContent(w http.ResponseWriter, r *http.Request) {
	util.LogGetRequest("/settings/glucose", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "settings/glucose.html", nil)
	}
}

// UserSettingsContent loads the content for user settings
func UserSettingsContent(w http.ResponseWriter, r *http.Request) {
	util.LogGetRequest("/settings/user", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "settings/user.html", nil)
	}
}
