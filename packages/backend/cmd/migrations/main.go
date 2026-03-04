package main

import (
	"log"

	"github.com/0xlebogang/shapas/internal/config"
	"github.com/0xlebogang/shapas/internal/database"
	"github.com/0xlebogang/shapas/internal/domain/user"
)

var models = []interface{}{
	user.User{},
}

func main() {
	dbUrl := config.LoadConfig().DatabaseUrl

	db, err := database.Connect(dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close(db)

	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	log.Println("Migrations run successfully")
}
