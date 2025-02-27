package database

import (
	db_postgresSQL "cerberus/internal/database/postgresSQL"
	db_redis "cerberus/internal/database/redis"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

// Databases struct holds connections to different database systems.
// Currently, it only contains a connection to a PostgreSQL database.
type Databases struct {
	Postgres *gorm.DB
	Redis    *db_redis.RedisPack
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
func InitDatabases(p_config *config.ConfigData) (*Databases, error) {
	pdb, err := db_postgresSQL.ConnectPostgres(p_config)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to connect to postgresSQL: %s", err.Error()), logger.ERROR)
		return nil, err
	}

	rdb, err := db_redis.ConnectRedis(p_config)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to connect to redis: %s", err.Error()), logger.ERROR)
		return nil, err
	}

	return &Databases{
		Postgres: pdb,
		Redis:    rdb,
	}, nil
}
