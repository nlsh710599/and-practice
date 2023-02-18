package database

import (
	"log"

	"github.com/nlsh710599/and-practice/internal/database/model"

	"gorm.io/gorm"
)

type RDS interface {
	CreateTable() error
	CreateNumber(string, string) error
	ReadNumber(string) string
	UpdateNumber(string, string) error
	DeleteNumber(string) error
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

func (pg *postgresClient) CreateNumber(name string, value string) error {
	return pg.client.Create(&model.Number{Name: name, Value: value}).Error
}

func (pg *postgresClient) ReadNumber(valueName string) string {
	var res *model.Number
	if err := pg.client.Model(&model.Number{}).Where("name = ?", valueName).Find(&res).Error; err != nil {
		return ""
	}
	return res.Value
}

func (pg *postgresClient) UpdateNumber(name string, value string) error {
	return pg.client.Model(&model.Number{}).
		Where("name = ?", name).
		Update("value", value).Error
}

func (pg *postgresClient) DeleteNumber(name string) error {
	return pg.client.Where("name =  ?", name).Delete(&model.Number{}).Error
}
