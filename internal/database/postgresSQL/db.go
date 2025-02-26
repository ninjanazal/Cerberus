package db_postgresSQL

import (
	postgres_models "cerberus/internal/database/postgresSQL/models"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// region Public

// ConnectPostgres establishes a connection to a PostgreSQL database and performs initial setup.
//
// Parameters:
//   - p_cfg: A pointer to a ConfigData struct containing database configuration information.
//
// Returns:
//   - *gorm.DB: A pointer to the initialized database connection.
//   - error: An error if the connection or migration fails, nil otherwise.
//
// If any step fails, it logs an error message and returns nil with the corresponding error.
func ConnectPostgres(p_cfg *config.ConfigData) (*gorm.DB, error) {
	if p_cfg.PostgresData.Host == "" {
		logger.Log("PostgresSql database url not defined", logger.ERROR)
		return nil, fmt.Errorf("PostgresSql database url not defined")
	}

	dsn := p_cfg.PostgresData.GetDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log(fmt.Sprintf("Something went wrong - %s", err.Error()), logger.ERROR)
		return nil, err
	}

	logger.Log("🧲 Connected to postgresSQL successfully", logger.INFO)

	if err := migrate(db); err != nil {
		logger.Log(fmt.Sprintf("Failed to apply migrations: %s", err.Error()), logger.ERROR)
		return nil, err
	}

	if err := db.AutoMigrate(&postgres_models.User{}); err != nil {
		logger.Log(fmt.Sprintf("AutoMigration failed - %s", err.Error()), logger.ERROR)
		return nil, err
	}

	logger.Log("🌲 PostgresSQL migrations completed", logger.INFO)
	return db, nil
}

// endregion Public
// region Private

// migrate creates the "uuid-ossp" extension in the database if it doesn't already exist.
//
// This function executes a SQL command to create the "uuid-ossp" extension, which
// provides UUID generation functions in PostgreSQL.
//
// Parameters:
//   - p_db: A pointer to a gorm.DB instance representing the database connection.
//
// Returns:
//   - error: An error if the execution fails, or nil if successful.
func migrate(p_db *gorm.DB) error {
	return p_db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
}

// endregion Private
