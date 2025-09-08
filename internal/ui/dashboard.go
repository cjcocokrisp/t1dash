package ui

import (
	"io"
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/util"

	log "github.com/sirupsen/logrus"
)

// IndexTestPage is the handler for a test page
func IndexTestPage(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://127.0.0.1:8080/api/rand")
	if err != nil {
		log.Fatal("Failed to make http request")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body")
	}

	fields := map[string]string{
		"EGV": string(data),
	}

	util.LogGetRequest("/", r.RemoteAddr)
	templates.Templates.ExecuteTemplate(w, "index.html", fields)
}
