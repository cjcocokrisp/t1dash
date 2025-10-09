package ui

/* login.go handles the ui for both first time set up and logging in
 * after first time set up. */

import (
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

// WelcomePage is the handler for the welcome page
func WelcomePage(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if there are users in the table if there are redirect to login

	util.LogGetRequest("/welcome", r.RemoteAddr)
	//data := map[string]string{
	//	"content": "welcome/welcome.html",
	//}
	templates.Templates.ExecuteTemplate(w, "welcome/layout.html", nil)
}

func SetUpContent(w http.ResponseWriter, r *http.Request) {
	// TODO: REFACTOR FOR AN ERROR PAGE REDIRECT!
	util.LogGetRequest("/setup", r.RemoteAddr)
	if util.ValidateHTMXRequest(r) {
		templates.Templates.ExecuteTemplate(w, "welcome/setup.html", nil)
	}
}
