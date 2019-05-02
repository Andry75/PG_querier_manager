package db_migrator

import "log"

const createDbSql = `
DO
    $do$
        BEGIN
            CREATE EXTENSION IF NOT EXISTS dblink; -- enable extension
            IF EXISTS(SELECT 1 FROM pg_database WHERE datname = 'querier_manager') THEN
                RAISE NOTICE 'Database already exists';
            ELSE
                PERFORM dblink_exec('dbname=' || current_database() -- current db
                            , 'CREATE DATABASE querier_manager');
            END IF;
        END
        $do$;
`

func createDB() {
	db := createConnectionWithoutDB()
	_, err := db.Connection.Exec(createDbSql)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Disconnect()
		if err != nil {
			log.Fatal(err)
		}
	}()

}
