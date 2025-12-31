package config

// Struct for application configuration
type T1DashConfig struct {
	ServerPort     int    // Port where server runs on
	ServerHostname string // Hostname where server is run
	Insecure       bool   // If true do not use TLS
	TLSCACertPath  string // Path to ca cert for tls
	TLSCAKeyPath   string // Path to ca key for tls
	TLSCertPath    string // Path to cert for tls
	TLSKeyPath     string // Path to key for tls
	DBHostname     string // Host for DB
	DBPort         int    // Port for DB
	DBRootPassword string // Password for root postgres user
	DBDatabase     string // Database name for DB
	DBUser         string // User for DB
	DBPassword     string // Password for DB
	SessionTTL     int    // Session total time to live (hours)
	SessionTimeout int    // Session timeout (minutes)
}

// AppCfg is the global app config
var AppCfg T1DashConfig
