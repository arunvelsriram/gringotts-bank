package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func NewConnection(dbConnUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbConnUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, err

	}

	return db, nil
}
