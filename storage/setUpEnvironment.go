package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //sqliteDriver
)

func CreateDbSchemaIfNotExists(connectionString string) error {
	connection, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return err
	}
	defer connection.Close()

	err = connection.Ping()
	if err != nil {
		return err
	}

	queryTableExists, err := connection.Prepare("SELECT 1 FROM sqlite_master WHERE type = 'table' AND name=?")
	if err != nil {
		return err
	}
	defer queryTableExists.Close()

	var tableExists bool
	err = queryTableExists.QueryRow("Temperature").Scan(&tableExists)
	if err != nil && sql.ErrNoRows != err {
		return err
	}

	if !tableExists {
		createTable, err := connection.Prepare("CREATE TABLE Temperature (Value REAL NOT NULL, DateTime DATETIME DEFAULT CURRENT_TIMESTAMP)")
		if err != nil {
			return err
		}
		defer createTable.Close()
		createTable.Exec()
		if err != nil {
			return err
		}

	}

	queryIndexExists, err := connection.Prepare("SELECT 1 FROM sqlite_master WHERE type = 'index' AND name=?")
	if err != nil {
		return err
	}
	defer queryIndexExists.Close()

	var indexExists bool
	err = queryIndexExists.QueryRow("Temperature_DateTime_INDEX").Scan(&indexExists)
	if err != nil && sql.ErrNoRows != err {
		return err
	}

	if !indexExists {
		createIndex, err := connection.Prepare("CREATE INDEX Temperature_DateTime_INDEX ON Temperature (DateTime);")
		if err != nil {
			return err
		}
		defer createIndex.Close()
		createIndex.Exec()
		if err != nil {
			return err
		}

	}

	return nil
}
