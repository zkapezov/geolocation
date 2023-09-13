package database

import (
	"errors"
	"net"
	"strconv"

	"github.com/biter777/countries"
)

type GeoLocation struct {
	IPAddress    string
	CountryCode  string
	Country      string
	City         string
	Latitude     float64
	Longitude    float64
	MysteryValue int64
}

type GeoLocationInput struct {
	IPAddress    string
	CountryCode  string
	Country      string
	City         string
	Latitude     string
	Longitude    string
	MysteryValue string
}

func (input *GeoLocationInput) Validate() error {
	if net.ParseIP(input.IPAddress) == nil {
		return errors.New("invalid IP address")
	}
	country := countries.ByName(input.CountryCode)
	if country == countries.Unknown {
		return errors.New("invalid country code")
	}
	country = countries.ByName(input.Country)
	if country == countries.Unknown {
		return errors.New("invalid country name")
	}
	if input.City == "" {
		return errors.New("missing city name")
	}
	latitude, err := strconv.ParseFloat(input.Latitude, 64)
	if err != nil {
		return errors.New("invalid latitude")
	}
	if latitude < -90 || latitude > 90 {
		return errors.New("incorrect latitude")
	}
	longitude, err := strconv.ParseFloat(input.Longitude, 64)
	if err != nil {
		return errors.New("invalid longitude")
	}
	if longitude < -180 || longitude > 180 {
		return errors.New("incorrect longitude")
	}
	_, err = strconv.ParseInt(input.MysteryValue, 10, 64)
	if err != nil {
		return errors.New("incorrect mystery value")
	}
	return nil
}

func (input *GeoLocationInput) ConvertToGeolocation() (GeoLocation, error) {
	latitude, err := strconv.ParseFloat(input.Latitude, 64)
	if err != nil {
		return GeoLocation{}, errors.New("invalid latitude")
	}
	longitude, err := strconv.ParseFloat(input.Longitude, 64)
	if err != nil {
		return GeoLocation{}, errors.New("invalid longitude")
	}
	mysteryValue, err := strconv.ParseInt(input.MysteryValue, 10, 64)
	if err != nil {
		return GeoLocation{}, errors.New("incorrect mystery value")
	}
	return GeoLocation{
		IPAddress:    input.IPAddress,
		CountryCode:  input.CountryCode,
		Country:      input.Country,
		City:         input.City,
		Latitude:     latitude,
		Longitude:    longitude,
		MysteryValue: mysteryValue,
	}, nil
}
