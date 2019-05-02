package db_migrator

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Db struct {
	Connection *sql.DB
}

type ConnectionString interface {
	GetConnectionStringWithoutDB() string
	GetConnectionString() string
}

func (db *Db) Connect(connectionString ConnectionString) error {
	db_, err := sql.Open("postgres", connectionString.GetConnectionString())
	if err != nil {
		return err
	}
	db.Connection = db_
	return nil
}

func (db *Db) ConnectWithoutDB(connectionString ConnectionString) error {
	db_, err := sql.Open("postgres", connectionString.GetConnectionStringWithoutDB())
	if err != nil {
		return err
	}
	db.Connection = db_
	return nil
}

func (db *Db) Disconnect() error {
	return db.Connection.Close()
}

func (db Db) Query(query string) (*sql.Rows, error) {
	rows, err := db.Connection.Query(query)

	if err != nil {
		return nil, err
	}
	return rows, nil
}
