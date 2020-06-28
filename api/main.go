package main

import "2020/data"

func main() {
	for _, ds := range data.GetDataSources() {
		ds.LoadData()
		ds.RegisterEndpoints()
	}
}
