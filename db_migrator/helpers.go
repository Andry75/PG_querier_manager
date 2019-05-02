package db_migrator

import (
	"github.com/Andry75/PG_querier_manager/config_loader"
	"log"
)

func createConnectionWithoutDB() Db {
	db := Db{}
	err := db.ConnectWithoutDB(getConfigs())
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func createConnection() Db {
	db := Db{}
	err := db.Connect(getConfigs())
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func getConfigs() ConnectionString {
	return config_loader.Load()
}
