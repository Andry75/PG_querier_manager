package pg_adapter

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	Connection *sqlx.DB
}

type ConnectionString interface {
	GetConnectionString() string
}

func (db *Db) Connect(connectionString ConnectionString) error {
	db_, err := sqlx.Connect("postgres", connectionString.GetConnectionString())
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
