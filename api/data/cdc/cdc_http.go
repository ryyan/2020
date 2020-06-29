package cdc

import (
	"encoding/json"
	"net/http"
)

func (ds *cdcDataSource) RegisterEndpoints() {
	http.HandleFunc("/data/cdc/weeklydeaths", ds.weeklyDeathsHandler)
}

func (ds *cdcDataSource) weeklyDeathsHandler(w http.ResponseWriter, req *http.Request) {
	response, _ := json.Marshal(ds.WeeklyDeaths)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
