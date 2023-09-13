package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zkapezov/geolocation/pkg/api"
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

	api := api.NewAPI(dbHandler)

	r := mux.NewRouter()
	r.HandleFunc("/geolocations/{ipaddress}", api.GetGeoLocationByIPAddressHandler).Methods("GET")

	http.Handle("/", r)
	logrus.Info("API has started...")
	http.ListenAndServe(":8080", nil)
}
