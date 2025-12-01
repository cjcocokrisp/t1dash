package ui

import (
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/db"
	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

// DashboardPage is the handler for bringing you to the home of the dashboard
func DashboardPage(w http.ResponseWriter, r *http.Request) {
	// TODO: Add logic for sign and such
	if exists, err := db.CheckIfUsersExist(); !exists && err == nil {
		util.LogRedirect("/dashboard", "/welcome", r.RemoteAddr, "no users exists")
		http.Redirect(w, r, "/welcome", http.StatusMovedPermanently)
		return
	}

	util.LogGetRequest("/dashboard", r.RemoteAddr)
	templates.Templates.ExecuteTemplate(w, "dashboard/base.html", nil)
}

// DashboardContent fills in the main content of the dashboard page with content for it
func DashboardContent(w http.ResponseWriter, r *http.Request) {
	// TODO: logic for data charts and such will be added here
	util.LogGetRequest("/dashboard/dashboard", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "dashboard/dashboard.html", nil)
	} else {
		util.LogRedirect("/dashboard/dashboard", "/dashboard", r.RemoteAddr, "not an htmx request")
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
	}
}

// ReportsContent fills in the main content of the dashboard page with content for it
func ReportsContent(w http.ResponseWriter, r *http.Request) {
	// TODO: do logic and content for this
	util.LogGetRequest("/dashboard/reports", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "dashboard/reports.html", nil)
	} else {
		util.LogRedirect("/dashboard/reports", "/dashboard", r.RemoteAddr, "not an htmx request")
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
	}
}

// UploadContent fills in the main content of the dashboard page with content for it
func UploadContent(w http.ResponseWriter, r *http.Request) {
	// TODO: add content for this
	util.LogGetRequest("/dashboard/upload", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "dashboard/upload.html", nil)
	} else {
		util.LogRedirect("/dashboard/upload", "/dashboard", r.RemoteAddr, "not an htmx request")
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
	}
}
