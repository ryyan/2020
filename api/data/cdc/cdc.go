package cdc

type cdcDataSource struct {
	WeeklyDeaths []WeeklyDeath
}

type WeeklyDeath struct {
	WeekEndingDate string `json:"week_ending_date"`
	AgeGroup       string `json:"age_group"`
	CovidDeaths    int64  `json:"covid_19_deaths,string"`
	TotalDeaths    int64  `json:"total_deaths,string"`
}

func NewCdcDataSource() *cdcDataSource {
	return &cdcDataSource{}
}
