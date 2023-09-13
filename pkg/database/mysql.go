package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLHandler(dsn string) (*DatabaseHandler, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return NewDatabaseHandler(db), nil
}
