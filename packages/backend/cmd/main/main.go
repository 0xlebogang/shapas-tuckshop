package main

import (
	"log"

	"github.com/0xlebogang/shapas/backend/internal/config"
	"github.com/0xlebogang/shapas/backend/internal/database"
	"github.com/0xlebogang/shapas/backend/internal/server"
)

func main() {
	config := config.LoadEnv()

	db, err := database.Connect(config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close(db)

	server := server.New(config, db)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
