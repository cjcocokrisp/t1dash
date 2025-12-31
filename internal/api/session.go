package api

import (
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/db"
	"github.com/cjcocokrisp/t1dash/internal/session"
	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/crypto"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	util.LogPostRequest("/login", r.RemoteAddr)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Got an error parsing form", http.StatusInternalServerError)
		return
	}

	user, err := db.GetUserByUsername(r.FormValue("username"))
	if err != nil {
		//http.Error(w, "Error when getting user", http.StatusInternalServerError)
		templates.Templates.ExecuteTemplate(w, "login/username.html", nil)
		return
	}

	verified, err := crypto.VerifyPassword(user.Password, r.FormValue("password"))
	if !verified {
		//http.Error(w, "Invalid login attempt", http.StatusUnauthorized)
		templates.Templates.ExecuteTemplate(w, "login/password.html", nil)
		return
	}

	cookie, err := session.CreateNewSessionCookie(user.Id, r.RemoteAddr)
	if err != nil {
		http.Error(w, "Error creating cookie", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, cookie)

	w.Header().Set("HX-Redirect", "/dashboard")
}
