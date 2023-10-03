package database

import (
	"article/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=article port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	// migration
	db.AutoMigrate(
		&domain.User{},
		&domain.Post{},
	)

	return db, nil
}
