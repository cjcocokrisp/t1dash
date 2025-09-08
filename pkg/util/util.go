package util

import (
	log "github.com/sirupsen/logrus"
)

func LogGetRequest(path string, client string) {
	log.WithFields(log.Fields{
		"client": client,
		"path":   path,
	}).Info("GET Request")
}
