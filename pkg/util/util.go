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

func LogError(kind string, location string, err error) {
	log.WithFields(log.Fields{
		"location": location,
		"error":    err,
	}).Error(kind)
}

func ValidateHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
