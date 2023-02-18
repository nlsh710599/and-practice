package database

import (
	"log"

	"github.com/nlsh710599/and-practice/internal/database/model"

	"gorm.io/gorm"
)

type RDS interface {
	CreateTable() error
}

type postgresClient struct {
	client *gorm.DB
}

func (pg *postgresClient) CreateTable() error {
	if err := pg.client.Migrator().AutoMigrate(
		&model.Number{},
	); err != nil {
		log.Printf("Failed to migrate tables: %v", err)
		return err
	}

	return nil
}
