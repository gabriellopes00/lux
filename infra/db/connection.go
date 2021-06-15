package db

import (
	"fmt"
	"helpy/config"
	"helpy/pkg/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormPG struct{}

func (orm *GormPG) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASS,
		config.DB_NAME,
		config.DB_PORT,
		config.DB_SSL_MODE,
		config.DB_TIME_ZONE,
	)
	fmt.Println(dsn)
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
