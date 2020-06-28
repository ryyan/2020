package data

import (
	"2020/data/cdc"
)

type DataSource interface {
	LoadData()
	RegisterEndpoints()
}

type EndpointHandler interface {
	
}

func GetDataSources() []DataSource {
	return []DataSource {
		cdc.NewCdcDataSource(),
	}
}
