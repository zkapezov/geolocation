package parser

import (
	"encoding/csv"
	"os"

	"github.com/zkapezov/geolocation/pkg/database"
)

type Parser struct {
	dbh *database.DatabaseHandler
}

func NewParser(dbh *database.DatabaseHandler) *Parser {
	return &Parser{dbh: dbh}
}

func (s *Parser) ParseCSVFile(filename string) ([][]string, error) {
	var records [][]string = make([][]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		return records, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err = reader.ReadAll()
	if err != nil {
		return records, err
	}

	return records, nil
}

func (s *Parser) SanitizeData(records [][]string) []database.GeoLocation {
	locations := make([]database.GeoLocation, 0)
	ipAddress := make(map[string]bool)
	for _, record := range records {
		if len(record) < 7 {
			continue
		}
		input := database.GeoLocationInput{
			IPAddress:    record[0],
			CountryCode:  record[1],
			Country:      record[2],
			City:         record[3],
			Latitude:     record[4],
			Longitude:    record[5],
			MysteryValue: record[6],
		}
		err := input.Validate()
		if err != nil {
			continue
		}
		location, err := input.ConvertToGeolocation()
		if err != nil {
			continue
		}
		if _, exists := ipAddress[location.IPAddress]; exists {
			continue
		}
		locations = append(locations, location)
	}
	return locations
}

func (s *Parser) PersistData(locations []database.GeoLocation) error {
	err := s.dbh.SaveGeoLocations(locations, true)
	return err
}
