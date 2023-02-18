package database

import (
	"errors"
	"log"

	"github.com/nlsh710599/and-practice/internal/database/model"

	"gorm.io/gorm"
)

var ErrNumberNotFound = errors.New("number does not exsit")

type RDS interface {
	CreateTable() error
	CreateNumber(string, string) error
	ReadNumber([]string) (map[string]string, error)
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

func (pg *postgresClient) ReadNumber(valueNames []string) (map[string]string, error) {
	var tmp *[]model.Number
	numberMap := make(map[string]string)

	if err := pg.client.Where("name IN ?", valueNames).Find(&tmp).Error; err != nil {
		return nil, err
	}

	for _, x := range *tmp {
		numberMap[x.Name] = x.Value
	}

	return numberMap, nil
}

func (pg *postgresClient) UpdateNumber(name string, value string) error {
	res := pg.client.Model(&model.Number{}).
		Where("name = ?", name).
		Update("value", value)

	if res.RowsAffected == 0 {
		return ErrNumberNotFound
	}
	return res.Error
}

func (pg *postgresClient) DeleteNumber(name string) error {
	res := pg.client.Where("name =  ?", name).Delete(&model.Number{})
	if res.RowsAffected == 0 {
		return ErrNumberNotFound
	}
	return res.Error
}
