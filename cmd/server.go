package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/api"
	"github.com/cjcocokrisp/t1dash/internal/config"
	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/internal/ui"
	"github.com/cjcocokrisp/t1dash/pkg/env"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// newServerCommand creates the server command for the cli
func newServerCommand() *cobra.Command {

	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Runs T1 Dash server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}

	serverCmd.Flags().IntVarP(&config.AppCfg.ServerPort, "port", "p", env.ParseNum("T1DASH_PORT", 8080, 0, 65535), "Port for the server to run on, defaults to 8080 and can also be set with the env variable T1DASH_PORT")
	serverCmd.Flags().StringVar(&config.AppCfg.ServerHostname, "host", env.ParseString("T1DASH_HOST", "localhost"), "Host for the server when it makes requests to the API")
	return serverCmd
}

// runServer runs the application and api server
func runServer() {
	templates.InitTemplates()

	r := chi.NewRouter()

	staticFS, err := fs.Sub(templates.WebFS, "web/static")
	if err != nil {
		log.Fatal(err)
	}
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	r.Get("/", ui.IndexTestPage)
	r.Get("/api/rand", api.GenerateRandomEGV)

	log.WithFields(log.Fields{
		"host": config.AppCfg.ServerHostname,
		"port": config.AppCfg.ServerPort,
	}).Info("T1 Dash Server Started")
	http.ListenAndServe(fmt.Sprintf(":%d", config.AppCfg.ServerPort), r)
}
