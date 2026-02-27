package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connectionString string) (*gorm.DB, error) {
	dialctor := postgres.Open(connectionString)
	return gorm.Open(dialctor, &gorm.Config{})
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get active database connection")
	}
	if err := sqlDB.Close(); err != nil {
		panic("Failed to close active database connection")
	}
}
