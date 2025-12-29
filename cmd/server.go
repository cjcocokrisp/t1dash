package main

import (
	"fmt"
	"io"
	"io/fs"
	stdlog "log"
	"net/http"

	"github.com/cjcocokrisp/t1dash/internal/api"
	"github.com/cjcocokrisp/t1dash/internal/config"
	"github.com/cjcocokrisp/t1dash/internal/db"
	"github.com/cjcocokrisp/t1dash/internal/templates"
	"github.com/cjcocokrisp/t1dash/internal/ui"
	"github.com/cjcocokrisp/t1dash/pkg/crypto"
	"github.com/cjcocokrisp/t1dash/pkg/env"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// newServerCommand creates the server command for the cli
func newServerCommand() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Runs T1 Dash server",
		Run: func(cmd *cobra.Command, args []string) {
			// Set up database then run server
			db.InitDBURL(config.AppCfg.DBHostname, config.AppCfg.DBPort, config.AppCfg.DBDatabase, config.AppCfg.DBUser, config.AppCfg.DBPassword)
			db.InitDBConnection()
			runServer()
		},
	}

	serverCmd.Flags().StringVar(&config.AppCfg.ServerHostname, "host", env.ParseString("T1DASH_HOST", "localhost"), "Host for the server when it makes requests to the API. Can also be set with T1DASH_HOST env variable.")
	serverCmd.Flags().IntVarP(&config.AppCfg.ServerPort, "port", "p", env.ParseNum("T1DASH_PORT", 8080, 0, 65535), "Port for the server to run on, defaults to 8080 and can also be set with the env variable T1DASH_PORT.")
	serverCmd.Flags().StringVar(&config.AppCfg.DBHostname, "db-host", env.ParseString("DB_HOST", "localhost"), "Hostname for Postgres DB. Can be set with DB_HOST.")
	serverCmd.Flags().IntVar(&config.AppCfg.DBPort, "db-port", env.ParseNum("DB_PORT", 5432, 0, 65535), "Port for the database, can be set with DB_PORT.")
	serverCmd.Flags().StringVar(&config.AppCfg.DBDatabase, "db-database", env.ParseString("DB_DATABASE", "t1dash"), "Name of the database, can be set with DB_DATABASE.")
	serverCmd.Flags().StringVar(&config.AppCfg.DBUser, "db-user", env.ParseString("DB_USER", "t1dash_user"), "Name of user for database, can be set with DB_USER.")
	serverCmd.Flags().StringVar(&config.AppCfg.DBPassword, "db-password", env.ParseString("DB_PASSWORD", "t1dash"), "Password of DB user, left empty if no value provided. can be set with DB_PASSWORD")
	serverCmd.Flags().StringVar(&config.AppCfg.DBRootPassword, "db-root-password", env.ParseString("DB_ROOT_PASSWORD", ""), "Password for the postgres root user. Used to create database and user. Can be set with DB_ROOT_PASSWORD")
	serverCmd.Flags().BoolVar(&config.AppCfg.Insecure, "insecure", false, "If set do not use TLS")
	serverCmd.Flags().StringVar(&config.AppCfg.TLSCACertPath, "tls-ca-cert-path", env.ParseString("TLS_CA_CERT_PATH", "ca.crt"), "Path to ca certificate for TLS, if it does not exist it will be created at that path")
	serverCmd.Flags().StringVar(&config.AppCfg.TLSCAKeyPath, "tls-ca-key-ath", env.ParseString("TLS_CA_KEY_PATH", "ca.key"), "Path to ca key for TLS, if it does not exist it will be created at that parth")
	serverCmd.Flags().StringVar(&config.AppCfg.TLSCertPath, "tls-cert-path", env.ParseString("TLS_CERT_PATH", "server.crt"), "Path to certificate for TLS, if it does not exist it will be created at that path")
	serverCmd.Flags().StringVar(&config.AppCfg.TLSKeyPath, "tls-key-path", env.ParseString("TLS_KEY_PATH", "server.key"), "Path to key for TLS, if it does not exist it will be created at that path")
	return serverCmd
}

// runServer runs the application and api server
func runServer() {
	// Init TLS
	if !config.AppCfg.Insecure {
		success := crypto.InitTLS()
		if success {
			log.Info("TLS cert and key exists, launching server")
		} else {
			log.Error("Error during TLS init, defaulting to http")
			config.AppCfg.Insecure = true
		}
	}

	if config.AppCfg.Insecure {
		log.Warn("TLS not configured, requests will be insecure!")
	}

	templates.InitTemplates()

	r := chi.NewRouter()

	staticFS, err := fs.Sub(templates.WebFS, "web/static")
	if err != nil {
		log.Fatal(err)
	}
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	r.Get("/dashboard", ui.DashboardPage)
	r.Get("/dashboard/dashboard", ui.DashboardContent)
	r.Get("/dashboard/reports", ui.ReportsContent)
	r.Get("/dashboard/upload", ui.UploadContent)
	r.Get("/settings/app", ui.AppSettingsContent)
	r.Get("/settings/glucose", ui.GlucoseSettingsContent)
	r.Get("/settings/user", ui.UserSettingsContent)
	r.Get("/login", ui.LoginPage)
	r.Get("/welcome", ui.SetupPage)

	r.Post("/welcome", api.InitialSetupAccountCreation)
	r.Post("/login", api.LoginUser)

	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", config.AppCfg.ServerPort),
		Handler:  r,
		ErrorLog: stdlog.New(io.Discard, "", 0), // Hide errors for TLS handshake, if needed will remove to debug
	}

	log.WithFields(log.Fields{
		"host": config.AppCfg.ServerHostname,
		"port": config.AppCfg.ServerPort,
	}).Info("T1 Dash Server Started")
	if config.AppCfg.Insecure {
		server.ListenAndServe()
	} else {
		server.ListenAndServeTLS("server.crt", "server.key")
	}
}
