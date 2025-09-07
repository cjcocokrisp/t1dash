package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/pkg/env"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

type ServerConfig struct {
	Port int
}

func newServerCommand() *cobra.Command {
	cfg := &ServerConfig{}

	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Runs T1 Dash server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer(cfg)
		},
	}

	serverCmd.Flags().IntVar(&cfg.Port, "port", env.ParseNum("T1DASH_PORT", 8080, 0, 65535), "Port for the server to run on, defaults to 8080")
	return serverCmd
}

func runServer(cfg *ServerConfig) {
	templates.InitTemplates()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		min := 50
		max := 350
		data := map[string]int{
			"EGV": rand.Intn(max-min+1) + min,
		}

		templates.Templates.ExecuteTemplate(w, "index.html", data)
	})

	fmt.Printf("Started http server on port :%d\n", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
}
