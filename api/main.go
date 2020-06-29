package main

import (
	"log"
	"net/http"

	"2020/data"
)

func main() {
	for _, ds := range data.GetDataSources() {
		ds.LoadData()
		ds.RegisterEndpoints()
	}

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
