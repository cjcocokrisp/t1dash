package main

import (
	"fmt"
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

func newServerCommand() *cobra.Command {

	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Runs T1 Dash server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}

	serverCmd.Flags().IntVar(&config.AppCfg.Server.Port, "port", env.ParseNum("T1DASH_PORT", 8080, 0, 65535), "Port for the server to run on, defaults to 8080")
	return serverCmd
}

func runServer() {
	templates.InitTemplates()

	r := chi.NewRouter()
	r.Get("/", ui.IndexTestPage)
	r.Get("/api/rand", api.GenerateRandomEGV)

	log.WithFields(log.Fields{
		"port": config.AppCfg.Server.Port,
	}).Info("T1 Dash Server Started")
	http.ListenAndServe(fmt.Sprintf(":%d", config.AppCfg.Server.Port), r)
}
