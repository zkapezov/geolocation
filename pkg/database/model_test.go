package database

import (
	"testing"
)

func TestGeoLocationInput_Validate(t *testing.T) {
	validInput := GeoLocationInput{
		IPAddress:    "192.168.1.1",
		CountryCode:  "US",
		Country:      "United States",
		City:         "New York",
		Latitude:     "40.7128",
		Longitude:    "-74.0060",
		MysteryValue: "12345",
	}

	if err := validInput.Validate(); err != nil {
		t.Errorf("Expected no validation error, but got: %v", err)
	}

	// Test invalid IP address
	invalidIPInput := GeoLocationInput{
		IPAddress:    "invalid_ip",
		CountryCode:  "US",
		Country:      "United States",
		City:         "New York",
		Latitude:     "40.7128",
		Longitude:    "-74.0060",
		MysteryValue: "12345",
	}
	if err := invalidIPInput.Validate(); err == nil {
		t.Error("Expected validation error for invalid IP address, but got none")
	}

	// Test invalid latitude
	invalidLatitudeInput := GeoLocationInput{
		IPAddress:    "192.168.1.1",
		CountryCode:  "US",
		Country:      "United States",
		City:         "New York",
		Latitude:     "invalid_latitude",
		Longitude:    "-74.0060",
		MysteryValue: "12345",
	}
	if err := invalidLatitudeInput.Validate(); err == nil {
		t.Error("Expected validation error for invalid latitude, but got none")
	}

	// Test invalid Longitude
	invalidLongitudeInput := GeoLocationInput{
		IPAddress:    "192.168.1.1",
		CountryCode:  "US",
		Country:      "United States",
		City:         "New York",
		Latitude:     "40.7128",
		Longitude:    "360",
		MysteryValue: "12345",
	}
	if err := invalidLongitudeInput.Validate(); err == nil {
		t.Error("Expected validation error for invalid longitude, but got none")
	}

	// Test invalid MysteryValue
	invalidMysteryInput := GeoLocationInput{
		IPAddress:    "192.168.1.1",
		CountryCode:  "US",
		Country:      "United States",
		City:         "New York",
		Latitude:     "40.7128",
		Longitude:    "-74.0060",
		MysteryValue: "invalid",
	}
	if err := invalidMysteryInput.Validate(); err == nil {
		t.Error("Expected validation error for invalid MysteryValue, but got none")
	}

	// Test missing country code
	missingCountyInput := GeoLocationInput{
		IPAddress:    "192.168.1.1",
		CountryCode:  "US",
		Country:      "",
		City:         "New York",
		Latitude:     "40.7128",
		Longitude:    "-74.0060",
		MysteryValue: "12345",
	}
	if err := missingCountyInput.Validate(); err == nil {
		t.Error("Expected validation error for missing country, but got none")
	}
}

func TestGeoLocationInput_ConvertToGeolocation(t *testing.T) {
	validInput := GeoLocationInput{
		IPAddress:    "192.168.1.1",
		CountryCode:  "US",
		Country:      "United States",
		City:         "New York",
		Latitude:     "40.7128",
		Longitude:    "-74.0060",
		MysteryValue: "12345",
	}

	geoLocation, err := validInput.ConvertToGeolocation()
	if err != nil {
		t.Errorf("Expected no conversion error, but got: %v", err)
	}

	if geoLocation.IPAddress != validInput.IPAddress {
		t.Errorf("Expected IP address %s, but got %s", validInput.IPAddress, geoLocation.IPAddress)
	}
	if geoLocation.CountryCode != validInput.CountryCode {
		t.Errorf("Expected CountryCode %s, but got %s", validInput.CountryCode, geoLocation.CountryCode)
	}
	if geoLocation.Country != validInput.Country {
		t.Errorf("Expected Country %s, but got %s", validInput.Country, geoLocation.Country)
	}
	if geoLocation.City != validInput.City {
		t.Errorf("Expected City %s, but got %s", validInput.City, geoLocation.City)
	}
	if geoLocation.Latitude != 40.7128 {
		t.Errorf("Expected Latitude %f, but got %f", 40.7128, geoLocation.Latitude)
	}
	if geoLocation.Longitude != -74.0060 {
		t.Errorf("Expected Longitude %f, but got %f", -74.0060, geoLocation.Longitude)
	}
	if geoLocation.MysteryValue != 12345 {
		t.Errorf("Expected IP address %d, but got %d", 12345, geoLocation.MysteryValue)
	}
}
