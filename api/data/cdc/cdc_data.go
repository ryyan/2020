package cdc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (ds *cdcDataSource) LoadData() {
	ds.getWeeklyDeaths()
}

func (ds *cdcDataSource) getWeeklyDeaths() {
	log.Println("Get weekly deaths")

	response, err := http.Get("https://data.cdc.gov/resource/vsak-wrfu.json?sex=All%20Sex")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []WeeklyDeath
	if err := json.Unmarshal(responseData, &data); err != nil {
		log.Fatal(err)
	}
	log.Println(data)

	ds.WeeklyDeaths = data
}
