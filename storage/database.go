package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase(connString string) (*Database, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection test failed: %w", err)
	}

	fmt.Println("Connected to database successfully")
	return &Database{Connection: db}, nil
}

func (d *Database) ExecuteQuery(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.Connection.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return result, nil
}

func (d *Database) Close() error {
	if err := d.Connection.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	fmt.Println("Database connection closed")
	return nil
}