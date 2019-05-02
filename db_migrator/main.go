package db_migrator

import (
	"github.com/DavidHuie/gomigrate"
	"log"
)

func Migrate() {
	createDB()
	db := createConnection()
	migrator, err := gomigrate.NewMigrator(db.Connection, gomigrate.Postgres{}, "./migrations")
	if err != nil {
		log.Fatal(err)
	}
	err = migrator.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = db.Disconnect()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
