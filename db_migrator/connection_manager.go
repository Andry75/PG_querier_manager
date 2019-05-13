package db_migrator

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Db struct {
	Connection *sql.DB
}

type ConnectionString interface {
	GetConnectionStringWithoutDB() (string, error)
	GetConnectionString() (string, error)
}

func (db *Db) Connect(connectionString ConnectionString) error {
	conStr, err := connectionString.GetConnectionString()
	if err != nil {
		return err
	}

	db_, err := sql.Open("postgres", conStr)
	if err != nil {
		return err
	}
	db.Connection = db_
	return nil
}

func (db *Db) ConnectWithoutDB(connectionString ConnectionString) error {
	conStr, err := connectionString.GetConnectionString()
	if err != nil {
		return err
	}

	db_, err := sql.Open("postgres", conStr)
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
