package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connectionString string) (*gorm.DB, error) {
	dialector := postgres.Open(connectionString)
	return gorm.Open(dialector, &gorm.Config{TranslateError: true})
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get active database connection")
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal("Failed close active database connection")
	}
}
