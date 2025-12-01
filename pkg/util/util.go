package util

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LogGetRequest(path string, client string) {
	log.WithFields(log.Fields{
		"client": client,
		"path":   path,
	}).Info("GET Request")
}

func LogRedirect(source, destination, client, reason string) {
	log.WithFields(log.Fields{
		"source":      source,
		"destination": destination,
		"client":      client,
		"reason":      reason,
	}).Info("Redirecting client")
}

func LogError(kind string, location string, err error) {
	log.WithFields(log.Fields{
		"location": location,
		"error":    err,
	}).Error(kind)
}

func LogPgError(code string, message string) {
	log.WithFields(log.Fields{
		"code":    code,
		"message": message,
	}).Error("DATABASE")
}

func ValidateHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
