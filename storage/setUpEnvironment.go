package storage

import "database/sql"

func CreateDbSchemaIfNotExists(connectionString string) error {
	connection, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return err
	}

	query, err := connection.Prepare("SELECT 1 FROM sqlite_master WHERE type = 'table' AND name=?")
	var exists bool
	query.QueryRow("Temperature").Scan(&exists)
	query.Close()

	if !exists {
		createTable, err := connection.Prepare("CREATE TABLE Temperature (Value REAL NOT NULL, DateTime DATETIME DEFAULT CURRENT_TIMESTAMP)")
		if err != nil {
			return err
		}
		createTable.Exec()
		createTable.Close()
	}
	return nil
}
