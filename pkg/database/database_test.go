package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDatabaseHandler_CreateTables(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	dh := NewDatabaseHandler(db)

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS geolocations").WillReturnResult(sqlmock.NewResult(0, 1))

	err = dh.CreateTables()
	if err != nil {
		t.Errorf("Error creating tables: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestDatabaseHandler_SaveGeoLocation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	dh := NewDatabaseHandler(db)

	location := GeoLocation{
		IPAddress:    "192.168.1.1",
		CountryCode:  "US",
		Country:      "United States",
		City:         "New York",
		Latitude:     40.7128,
		Longitude:    -74.0060,
		MysteryValue: 348,
	}

	mock.ExpectExec("INSERT INTO geolocations").WillReturnResult(sqlmock.NewResult(0, 1))

	err = dh.SaveGeoLocation(location)
	if err != nil {
		t.Errorf("Error saving geolocation: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestDatabaseHandler_SaveGeoLocations(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dh := NewDatabaseHandler(db)
	var locations []GeoLocation

	count := 10000
	for i := 0; i < count; i++ {
		locations = append(locations, GeoLocation{
			IPAddress:    "192.168.1.1",
			CountryCode:  "US",
			Country:      "United States",
			City:         "New York",
			Latitude:     40.7128,
			Longitude:    -74.0060,
			MysteryValue: 348,
		})
	}

	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))
	mock.ExpectPrepare("INSERT INTO geolocations").ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1000))

	err = dh.SaveGeoLocations(locations, false)
	if err != nil {
		t.Errorf("Error saving geolocation: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestDatabaseHandler_GetGeoLocationByIPAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dh := NewDatabaseHandler(db)

	// Define the expected input IP address and corresponding geolocation data
	expectedIPAddress := "192.168.1.1"
	expectedGeoLocation := GeoLocation{
		IPAddress:    expectedIPAddress,
		CountryCode:  "US",
		Country:      "United States",
		City:         "New York",
		Latitude:     40.7128,
		Longitude:    -74.0060,
		MysteryValue: 23445,
	}

	query := "SELECT ip_address, country_code, country, city, latitude, longitude, mystery_value FROM geolocations WHERE ip_address = ?"
	mock.ExpectQuery(query).
		WithArgs(expectedIPAddress).
		WillReturnRows(sqlmock.NewRows([]string{"ip_address", "country_code", "country", "city", "latitude", "longitude", "mystery_value"}).
			AddRow(
				expectedGeoLocation.IPAddress,
				expectedGeoLocation.CountryCode,
				expectedGeoLocation.Country,
				expectedGeoLocation.City,
				expectedGeoLocation.Latitude,
				expectedGeoLocation.Longitude,
				expectedGeoLocation.MysteryValue,
			))

	geoLocation, err := dh.GetGeoLocationByIPAddress(expectedIPAddress)
	if err != nil {
		t.Fatalf("Error getting geolocation by IP address: %v", err)
	}

	if geoLocation.IPAddress != expectedGeoLocation.IPAddress {
		t.Errorf("Expected IP address %s, but got %s", expectedGeoLocation.IPAddress, geoLocation.IPAddress)
	}
	if geoLocation.CountryCode != expectedGeoLocation.CountryCode {
		t.Errorf("Expected Country Code %s, but got %s", expectedGeoLocation.CountryCode, geoLocation.CountryCode)
	}
	if geoLocation.Country != expectedGeoLocation.Country {
		t.Errorf("Expected Country %s, but got %s", expectedGeoLocation.Country, geoLocation.Country)
	}
	if geoLocation.City != expectedGeoLocation.City {
		t.Errorf("Expected City %s, but got %s", expectedGeoLocation.City, geoLocation.City)
	}
	if geoLocation.Latitude != expectedGeoLocation.Latitude {
		t.Errorf("Expected Latitude %f, but got %f", expectedGeoLocation.Latitude, geoLocation.Latitude)
	}
	if geoLocation.Longitude != expectedGeoLocation.Longitude {
		t.Errorf("Expected Longitude %f, but got %f", expectedGeoLocation.Longitude, geoLocation.Longitude)
	}
	if geoLocation.MysteryValue != expectedGeoLocation.MysteryValue {
		t.Errorf("Expected MysteryValue %d, but got %d", expectedGeoLocation.MysteryValue, geoLocation.MysteryValue)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
