package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/zkapezov/geolocation/pkg/parser"
	"github.com/zkapezov/geolocation/pkg/utils"
)

func main() {
	dbHandler, err := utils.GetDatabaseHandler()
	if err != nil {
		logrus.WithError(err).Panic("can't open database connection")
	}
	defer dbHandler.Close()
	err = dbHandler.CreateTables()
	if err != nil {
		logrus.WithError(err).Panic("can't create necessary tables")
	}

	service := parser.NewParser(dbHandler)
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "ERROR: exactly one argument must be provided, CSV filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	records, err := service.ParseCSVFile(filename)
	if err != nil {
		logrus.WithError(err).Error("can't parse csv file")
		fmt.Fprintln(os.Stderr, "ERROR: can't parse csv file")
		os.Exit(1)
	}
	locations := service.SanitizeData(records)
	err = service.PersistData(locations)
	if err != nil {
		logrus.WithError(err).Error("can't persist data into database")
		fmt.Fprintln(os.Stderr, "ERROR: can't persist data into database")
		os.Exit(1)
	}

	fmt.Printf("Total entries: %d\n", len(records))
	fmt.Printf("Accepted entries: %d\n", len(locations))
	fmt.Printf("Discarded entries: %d\n", len(records)-len(locations))
}
