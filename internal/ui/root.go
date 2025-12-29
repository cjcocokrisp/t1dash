package ui

import (
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/db"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

func RootRedirects(w http.ResponseWriter, r *http.Request) {
	if exists, err := db.CheckIfUsersExist(); !exists && err == nil {
		util.LogRedirect("/root", "/welcome", r.RemoteAddr, "no users exists")
		http.Redirect(w, r, "/welcome", http.StatusMovedPermanently)
		return
	}

	// TODO: check sessions
	util.LogRedirect("/root", "/welcome", r.RemoteAddr, "no session")
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}
