package db_postgresSQL

import (
	postgres_models "cerberus/internal/database/postgresSQL/models"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	//Run migrations
	if err := db.AutoMigrate(&postgres_models.User{}); err != nil {
		logger.Log(fmt.Sprintf("Migration failed - %s", err.Error()), logger.ERROR)
		return nil, err
	}

	logger.Log("🌲 Migrations completed, postgres setup ended", logger.INFO)

	return db, nil
}
