package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type DatabaseHandler struct {
	db *sql.DB
}

func NewDatabaseHandler(db *sql.DB) *DatabaseHandler {
	return &DatabaseHandler{db: db}
}

func (dh *DatabaseHandler) CreateTables() error {
	_, err := dh.db.Exec(
		`CREATE TABLE IF NOT EXISTS geolocations (
			id INT AUTO_INCREMENT PRIMARY KEY,
			ip_address VARCHAR(255) NOT NULL,
			country_code VARCHAR(255),
			country VARCHAR(255),
			city VARCHAR(255),
			latitude DECIMAL(10, 8),
			longitude DECIMAL(11, 8),
			mystery_value BIGINT,

			INDEX (ip_address)
		);`,
	)
	return err
}

func (dh *DatabaseHandler) SaveGeoLocation(location GeoLocation) error {
	_, err := dh.db.Exec(
		"INSERT INTO geolocations (ip_address, country_code, country, city, latitude, longitude, mystery_value) VALUES (?, ?, ?, ?, ?, ?, ?)",
		location.IPAddress, location.CountryCode, location.Country, location.City, location.Latitude, location.Longitude, location.MysteryValue,
	)
	return err
}

func (dh *DatabaseHandler) saveGeoLocationChunk(locations []GeoLocation) error {
	query := "INSERT INTO geolocations (ip_address, country_code, country, city, latitude, longitude, mystery_value) VALUES "
	var inserts []string
	var params []interface{}
	for _, location := range locations {
		inserts = append(inserts, "(?, ?, ?, ?, ?, ?, ?)")
		params = append(params,
			location.IPAddress,
			location.CountryCode,
			location.Country,
			location.City,
			location.Latitude,
			location.Longitude,
			location.MysteryValue,
		)
	}
	query = query + strings.Join(inserts, ",")
	stmt, err := dh.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(params...)
	return err
}

func (dh *DatabaseHandler) SaveGeoLocations(locations []GeoLocation, outputProcess bool) error {
	chunkSize := 1000
	for i := 0; i < len(locations); i += chunkSize {
		end := i + chunkSize
		if end > len(locations) {
			end = len(locations)
		}
		locationsToInsert := locations[i:end]
		err := dh.saveGeoLocationChunk(locationsToInsert)
		if err != nil {
			return err
		}
		if outputProcess {
			fmt.Printf("Inserted %d out of %d locations\n", end, len(locations))
		}
	}
	return nil
}

func (dh *DatabaseHandler) GetGeoLocationByIPAddress(ipAddress string) (GeoLocation, error) {
	row := dh.db.QueryRow("SELECT ip_address, country_code, country, city, latitude, longitude, mystery_value FROM geolocations WHERE ip_address = ?", ipAddress)
	var location GeoLocation
	err := row.Scan(
		&location.IPAddress,
		&location.CountryCode,
		&location.Country,
		&location.City,
		&location.Latitude,
		&location.Longitude,
		&location.MysteryValue,
	)
	return location, err
}

func (dh *DatabaseHandler) Close() error {
	return dh.db.Close()
}
