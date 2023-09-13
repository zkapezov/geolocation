package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteHandler(dbPath string) (*DatabaseHandler, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return NewDatabaseHandler(db), nil
}
