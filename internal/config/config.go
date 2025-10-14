package config

// Struct for application configuration
type T1DashConfig struct {
	ServerPort     int    // Port where server runs on
	ServerHostname string // Hostname where server is run
	DBHostname     string // Host for DB
	DBPort         int    // Port for DB
	DBRootPassword string // Password for root postgres user
	DBDatabase     string // Database name for DB
	DBUser         string // User for DB
	DBPassword     string // Password for DB
}

// Global app config
var AppCfg T1DashConfig
