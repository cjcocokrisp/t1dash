package session

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cjcocokrisp/t1dash/internal/config"
	"github.com/cjcocokrisp/t1dash/internal/db"
	"github.com/cjcocokrisp/t1dash/pkg/util"

	"github.com/jackc/pgx/v5/pgtype"
)

const cookieName = "t1DashSession"

// ValidateSession checks to see if the session is valid will return true if it is valid
// A session is expired with the following three conditions
// 1) Past Expires at time
// 2) The last time is was seen was over time specified
// 3) Session valid was set to false by some other function
func ValidateSession(sessionId pgtype.UUID) (bool, error) {
	session, err := db.GetSessionById(sessionId)
	if err != nil {
		return false, err
	}

	now := time.Now()
	if now.After(session.ExpiresAt) {
		db.InvalidateSession(sessionId)
		return false, nil
	}

	inactive := session.LastSeen.Add(time.Duration(config.AppCfg.SessionTimeout) * time.Minute)
	fmt.Println(time.Duration(config.AppCfg.SessionTimeout) * time.Minute)
	if now.After(inactive) {
		db.InvalidateSession(sessionId)
		return false, nil
	}

	return session.Valid, nil
}

// CreateNewSessionCoookie creates a new cookie for a session
func CreateNewSessionCookie(userId pgtype.UUID, ip string) (*http.Cookie, error) {
	sessionId, err := db.CreateSession(userId, ip)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    sessionId.String(),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}
	return cookie, nil
}

// ParseSessionCookie parses the an http request for the session cookie
func ParseSessionCookie(r *http.Request) (*pgtype.UUID, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	var sessionId pgtype.UUID
	err = sessionId.Scan(cookie.Value)
	if err != nil {
		return nil, err
	}

	return &sessionId, nil
}

// parseAndValidateSessionCookie is a helper function for the RedirectOn functions
// It parses the session cookie and then validates it and returns if it is valid or not along
// with if there was an error, the first bool is if its valid and the second is if there was an error
// It also updates the last seen value for the session if its valid
func parseAndValidateSessionCookie(w *http.ResponseWriter, r *http.Request, currentPath string) (bool, bool) {
	sessionId, err := ParseSessionCookie(r)
	if err != nil {
		if !errors.Is(err, http.ErrNoCookie) {
			util.LogError("session", currentPath, err)
			http.Error(*w, "Error parsing session cookie that was not it being found", http.StatusInternalServerError)
			return false, true
		}
	}

	valid, err := ValidateSession(*sessionId)
	if err != nil {
		util.LogError("session", currentPath, err)
		http.Error(*w, "Error validating session", http.StatusInternalServerError)
		return false, true
	}

	if valid {
		err = db.UpdateLastSeen(*sessionId, time.Now())
		if err != nil {
			util.LogError("session", currentPath, err)
			http.Error(*w, "Error updating last seen", http.StatusInternalServerError)
			return false, true
		}
	}

	return valid, false
}

// RedirectOnValidSession parses an http request and then redirects if the session cookie
// is valid, will return true if redirected and false if it did not
func RedirectOnValidSession(w *http.ResponseWriter, r *http.Request, currentPath, newPath string) bool {
	valid, err := parseAndValidateSessionCookie(w, r, currentPath)
	if err {
		return false
	}

	if !valid {
		return false
	}

	util.LogRedirect(currentPath, newPath, r.RemoteAddr, "client has a valid session")
	http.Redirect(*w, r, newPath, http.StatusMovedPermanently)
	return true
}

// RedirectOnInvalidSession parses an http request and then redirects if the session cookie
// is not valid, will return true if it was redirected and false if it was not
func RedirectOnInvalidSession(w *http.ResponseWriter, r *http.Request, currentPath, newPath string) bool {
	valid, err := parseAndValidateSessionCookie(w, r, currentPath)
	if err {
		return false
	}

	if valid {
		return false
	}

	util.LogRedirect(currentPath, newPath, r.RemoteAddr, "client has a invalid session")
	http.Redirect(*w, r, newPath, http.StatusMovedPermanently)
	return true
}
