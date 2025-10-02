package config

// Struct for application configuration
type T1DashConfig struct {
	ServerPort     int    // Port where server runs on
	ServerHostname string // Hostname where server is run
}

// Global app config
var AppCfg T1DashConfig
