package db_config

import (
	"fmt"
	"os"
)

// PostgresConfigData represents the configuration data for a database connection.
//
// This struct holds the necessary information to establish a connection
// to a database server, typically PostgreSQL.
type PostgresConfigData struct {
	Host    string // The hostname or IP address of the database server
	Port    string // The port number on which the database server is listening
	DbName  string // The name of the database to connect to
	SslMode string // The SSL mode for the connection (e.g., "disable", "require", "verify-full")
}

// DefaultPostgresCfg is a global variable that holds default configuration values.
//
// It is initialized with preset values in the init() function and can be used
// as a fallback or starting point for database configurations.
var DefaultPostgresCfg PostgresConfigData

// init initializes the DefaultCfg with preset values.
//
// This function is automatically called when the package is imported.
// It sets up default values for the database configuration, which can be
// overridden later if needed.
func init() {
	DefaultPostgresCfg.Host = "localhost"
	DefaultPostgresCfg.Port = "8080"
	DefaultPostgresCfg.DbName = "db"
	DefaultPostgresCfg.SslMode = "disable"
}

// region Public

// ParseLineData updates the ConfigData struct based on a key-value pair.
//
// This function takes a key and its corresponding value as strings and updates
// the appropriate field in the ConfigData struct. It handles specific
// PostgreSQL configuration keys.
//
// Parameters:
//   - p_key: A string representing the configuration key.
//   - p_value: A string representing the value for the given key.
func (cfg *PostgresConfigData) ParseLineData(p_key string, p_value string) {
	fMap := map[string]*string{
		"POSTGRES_HOST":    &cfg.Host,
		"POSTGRES_PORT":    &cfg.Port,
		"POSTGRES_DBNAME":  &cfg.DbName,
		"POSTGRES_SSLMODE": &cfg.SslMode,
	}

	if f, ok := fMap[p_key]; ok {
		*f = p_value
		return
	}
}

// GetDsn generates and returns a PostgreSQL connection string (DSN) based on the current configuration.
//
// The function retrieves the username and password from environment variables and combines them
// with other configuration parameters to create a complete DSN string.
//
// Returns:
//   - string: A formatted DSN string suitable for establishing a PostgreSQL database connection.
//
// Environment Variables Used:
//   - POSTGRES_USERNAME: The PostgreSQL username.
//   - POSTGRES_PASSWORD: The PostgreSQL password.
func (cfg *PostgresConfigData) GetDsn() string {
	var usr string = os.Getenv("POSTGRES_USERNAME")
	var pwd string = os.Getenv("POSTGRES_PASSWORD")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, usr, pwd, cfg.DbName, cfg.Port, cfg.SslMode)
}

// endregion Public
