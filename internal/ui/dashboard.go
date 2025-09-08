package ui

import (
	"math/rand"
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/util"
)

// IndexTestPage is the handler for a test page
func IndexTestPage(w http.ResponseWriter, r *http.Request) {
	min := 50
	max := 350
	data := map[string]int{
		"EGV": rand.Intn(max-min+1) + min,
	}

	util.LogGetRequest("/", r.RemoteAddr)
	templates.Templates.ExecuteTemplate(w, "index.html", data)
}
