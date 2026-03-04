package main

import (
	"log"

	"github.com/0xlebogang/shapas/internal/config"
	"github.com/0xlebogang/shapas/internal/database"
	"github.com/0xlebogang/shapas/internal/server"
)

func main() {
	config := config.LoadConfig()

	db, err := database.Connect(config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close(db)

	svr := server.New(config, db)
	if err := svr.Start(); err != nil {
		log.Fatalf("Failed to start server")
	}
}
