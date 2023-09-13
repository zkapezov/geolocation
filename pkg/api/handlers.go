package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zkapezov/geolocation/pkg/database"
)

type API struct {
	dbh *database.DatabaseHandler
}

func NewAPI(dbh *database.DatabaseHandler) *API {
	return &API{dbh: dbh}
}

func (api *API) GetGeoLocationByIPAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipAddress := vars["ipaddress"]

	geoLocation, err := api.dbh.GetGeoLocationByIPAddress(ipAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "IP Address not found", http.StatusNotFound)
		} else {
			logrus.WithError(err).Error("can't get geolocation data from db")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(geoLocation)
	if err != nil {
		logrus.WithError(err).Error("can't convert geolocation to json")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
