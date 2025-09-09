package config

// Struct for application configuration
type T1DashConfig struct {
	Server struct {
		Port int // Port where server runs on
	}
}

// Global app config
var AppCfg T1DashConfig
