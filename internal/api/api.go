package api

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/cjcocokrisp/t1dash/pkg/util"
)

// GenerateRandomEGV is an API endpoint that generates a random number
// between 50 - 350. Will be removed in a future revision here for
// testing at the moment
func GenerateRandomEGV(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	min := 50
	max := 350
	rand := rand.Intn(max-min+1) + min

	util.LogGetRequest("/api/rand", r.RemoteAddr)
	w.Write([]byte(strconv.Itoa(rand)))
}
