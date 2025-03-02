package database

import (
	"cerberus/internal/tools/jwt"
	"cerberus/internal/tools/logger"
	"cerberus/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

// DataRefs struct holds connections to different database systems.
// Currently, it only contains a connection to a PostgreSQL database.
type DataRefs struct {
	Postgres *gorm.DB

	Redis  *RedisPack
	JWTGen *jwt.JWTGenerator

	ConfigData *config.ConfigData
}

// InitDatabases initializes database connections based on the provided configuration.
//
// This function attempts to establish a connection to a PostgreSQL database
// using the provided configuration. If successful, it returns a Databases
// struct containing the established connection.
//
// Parameters:
//   - p_config: A pointer to a config.ConfigData struct containing the necessary
//     configuration information for database connections.
//
// Returns:
//   - *Databases: A pointer to a Databases struct containing the initialized
//     database connections. Currently, this only includes a PostgreSQL connection.
//   - error: An error if the connection attempt fails, or nil if successful.
//
// If the connection to PostgreSQL fails, this function will log an error
// message using the logger package before returning the error.
func InitDatabases(p_config *config.ConfigData) (*DataRefs, error) {
	pdb, err := ConnectPostgres(p_config)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to connect to postgresSQL: %s", err.Error()), logger.ERROR)
		return nil, err
	}

	rdb, err := ConnectRedis(p_config)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to connect to redis: %s", err.Error()), logger.ERROR)
		return nil, err
	}

	return &DataRefs{
		Postgres: pdb,

		Redis:  rdb,
		JWTGen: jwt.NewJWTGenerator(p_config),

		ConfigData: p_config,
	}, nil
}
