package ui

import (
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

// DashboardPage is the handler for bringing you to the home of the dashboard
func DashboardPage(w http.ResponseWriter, r *http.Request) {
	// TODO: Add logic for sign and such
	util.LogGetRequest("/dashboard", r.RemoteAddr)
	templates.Templates.ExecuteTemplate(w, "dashboard/base.html", nil)
}

// DashboardContent fills in the main content of the dashboard page with content for it
func DashboardContent(w http.ResponseWriter, r *http.Request) {
	// TODO: logic for data charts and such will be added here
	util.LogGetRequest("/dashboard/dashboard", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "dashboard/dashboard.html", nil)
	}
}

// ReportsContent fills in the main content of the dashboard page with content for it
func ReportsContent(w http.ResponseWriter, r *http.Request) {
	// TODO: do logic and content for this
	util.LogGetRequest("/dashboard/reports", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "dashboard/reports.html", nil)
	}
}

// UploadContent fills in the main content of the dashboard page with content for it
func UploadContent(w http.ResponseWriter, r *http.Request) {
	// TODO: add content for this
	util.LogGetRequest("/dashboard/upload", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "dashboard/upload.html", nil)
	}
}
