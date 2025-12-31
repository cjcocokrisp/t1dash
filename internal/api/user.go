package api

import (
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/db"
	"github.com/cjcocokrisp/t1dash/internal/models"
	"github.com/cjcocokrisp/t1dash/internal/session"
	"github.com/cjcocokrisp/t1dash/pkg/crypto"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

func InitialSetupAccountCreation(w http.ResponseWriter, r *http.Request) {
	util.LogPostRequest("/welcome", r.RemoteAddr)

	if exists, err := db.CheckIfUsersExist(); !exists && err != nil {
		http.Error(w, "User exists should not be making a post request here", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Got an error parsing form", http.StatusInternalServerError)
		return
	}

	hash, err := crypto.HashPassword(r.FormValue("password"))
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username:  r.FormValue("username"),
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Password:  hash,
		Avatar:    "needtocomeupwithhowthosearegoingtobestoredlmao",
		Role:      "admin",
	}

	userId, err := db.CreateUser(&user)
	if err != nil && userId != nil {
		http.Error(w, "Error creating db entry for user", http.StatusInternalServerError)
		return
	}

	cookie, err := session.CreateNewSessionCookie(*userId, r.RemoteAddr)
	if err != nil {
		http.Error(w, "Error creating cookie", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, cookie)

	util.LogRedirect("/welcome", "/dashboard", r.RemoteAddr, "inital user created")
	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
}
