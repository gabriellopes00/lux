package db

import (
	"helpy/pkg/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormPG struct{}

func (orm *GormPG) Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=psql dbname=helpy port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = orm.AutoMigrate(db)
	if err != nil {
		return db, err
	}

	return db, nil
}

func (orm *GormPG) AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	return nil
}
