package db

import (
	"fmt"
	"helpy/config/env"
	"helpy/pkg/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormPG struct{}

func (orm *GormPG) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		env.DB_HOST,
		env.DB_USER,
		env.DB_PASS,
		env.DB_NAME,
		env.DB_PORT,
		env.DB_SSL_MODE,
		env.DB_TIME_ZONE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
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
