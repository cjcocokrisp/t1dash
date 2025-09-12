package config

// Struct for application configuration
type T1DashConfig struct {
	Server struct {
		Port     int    // Port where server runs on
		Hostname string // Hostname where server is run
	}
}

// Global app config
var AppCfg T1DashConfig
