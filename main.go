package main

import (
	"github.com/Andry75/PG_querier_manager/db_migrator"
	"github.com/Andry75/PG_querier_manager/web_server"
	"log"
)

func main() {
	log.Println("Migrating DB start:")
	// Migrate DB
	db_migrator.Migrate()

	log.Println("Migrating DB finished\n\n")

	log.Println("Starting service")

	// Start server
	web_server.Start()
}
