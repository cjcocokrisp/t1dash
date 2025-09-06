package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

// TODO: Add embed file system into a common variable

func newServerCommand() *cobra.Command {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Runs T1 Dash server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}
	return serverCmd
}

func runServer() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../web/templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		min := 50
		max := 350
		data := map[string]int{
			"EGV": rand.Intn(max-min+1) + min,
		}

		tmpl.Execute(w, data)
	})

	fmt.Println("Started http server on port :3000")
	http.ListenAndServe(":3000", r)
}
