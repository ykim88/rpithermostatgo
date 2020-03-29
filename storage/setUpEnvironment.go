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

	query, err := connection.Prepare("SELECT 1 FROM sqlite_master WHERE type = 'table' AND name=?")
	if err != nil {
		return err
	}
	defer query.Close()
	var exists bool
	err = query.QueryRow("Temperature").Scan(&exists)
	if err != nil && sql.ErrNoRows != err {
		return err
	}

	if !exists {
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
	return nil
}
