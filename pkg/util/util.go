package util

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LogGetRequest(path, client string) {
	log.WithFields(log.Fields{
		"client": client,
		"path":   path,
	}).Info("GET Request")
}

func LogPostRequest(path, client string) {
	log.WithFields(log.Fields{
		"client": client,
		"path":   path,
	})
}

func LogRedirect(source, destination, client, reason string) {
	log.WithFields(log.Fields{
		"client":      client,
		"source":      source,
		"destination": destination,
		"reason":      reason,
	}).Info("Redirecting client")
}

func LogError(kind, location string, err error) {
	log.WithFields(log.Fields{
		"location": location,
		"error":    err,
	}).Error(kind)
}

func LogPgError(code, message string) {
	log.WithFields(log.Fields{
		"code":    code,
		"message": message,
	}).Error("DATABASE")
}

func ValidateHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
