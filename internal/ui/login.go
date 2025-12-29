package ui

/* login.go handles the ui for both first time set up and logging in
 * after first time set up. */

import (
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/db"
	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	// TODO: add logic
	if exists, err := db.CheckIfUsersExist(); !exists && err == nil {
		util.LogRedirect("/login", "/welcome", r.RemoteAddr, "no users exists")
		http.Redirect(w, r, "/welcome", http.StatusMovedPermanently)
		return
	}

	util.LogGetRequest("/login", r.RemoteAddr)
	templates.Templates.ExecuteTemplate(w, "login/base.html", nil)
}

// SetupPage is the handler for the setup/welcome page will be redirected from base or login
// based on if you have any users
func SetupPage(w http.ResponseWriter, r *http.Request) {
	// TODO: refactor to use multiple endpoints, it's better practice
	if exists, err := db.CheckIfUsersExist(); exists && err == nil {
		util.LogRedirect("/welcome", "/login", r.RemoteAddr, "a users exists")
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		return
	}

	content := r.URL.Query().Get("content")
	contentTmpl := "setup/base.html"
	if util.ValidateHTMXRequest(r) && content == "agreement" {
		contentTmpl = "setup/agreement.html"
	} else if util.ValidateHTMXRequest(r) && content == "setup" {
		contentTmpl = "setup/setup.html"
	}

	util.LogGetRequest("/welcome", r.RemoteAddr)
	templates.Templates.ExecuteTemplate(w, contentTmpl, nil)
}
