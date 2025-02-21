package database

import "gorm.io/gorm"

type Databases struct {
	Postgres *gorm.DB
}

func InitDatabases() (*Databases, error) {
	return &Databases{}, nil
}
