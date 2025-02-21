package db_postgresSQL

import (
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(p_cfg *config.ConfigData) (*gorm.DB, error) {
	if p_cfg.PostgresURL == "" {
		logger.Log("PostgresSql database url not defined", logger.ERROR)
		return nil, fmt.Errorf("PostgresSql database url not defined")
	}

	db, err := gorm.Open(postgres.Open(p_cfg.PostgresURL), &gorm.Config{})
	if err != nil {
		logger.Log(fmt.Sprintf("Something went wrong - %s", err.Error()), logger.ERROR)
		return nil, err
	}

	logger.Log("Connected to postgresSQL successfully", logger.INFO)

	return db, nil
}
