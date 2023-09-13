package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/zkapezov/geolocation/pkg/database"
)

func GetDatabaseHandler() (*database.DatabaseHandler, error) {
	databaseType := os.Getenv(EnvDatabaseType)
	if databaseType == EnvDatabaseSQLite {
		filepath := os.Getenv(EnvSQLitePath)
		return database.NewSQLiteHandler(filepath)
	} else if databaseType == EnvDatabaseMysql {
		dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s",
			os.Getenv(EnvMysqlUser),
			os.Getenv(EnvMysqlPassword),
			os.Getenv(EnvMysqlHost),
			os.Getenv(EnvMysqlPort),
			os.Getenv(EnvMysqlDatabase),
		)
		return database.NewMySQLHandler(dsn)
	}
	return nil, errors.New("unknown database type")
}
