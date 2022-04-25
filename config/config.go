package config

import (
	"enterpret/models"
	"time"
)

//SourceConfig stores information regarding various sources
type SourceConfig struct {
	Name   string
	Params models.Params
}

//TenantConfig stores information regarding a tenant including the sources and query parameters
type TenantConfig struct {
	Name    string
	Sources []SourceConfig
}

var TC = []TenantConfig{
	{
		Name: "zoom.us",
		Sources: []SourceConfig{
			{
				Name: "discourse",
				Params: models.Params{
					Since:       getTime(2000, 1, 1, 1, 1, 1, 1, time.Local),
					Before:      getTime(2023, 1, 1, 1, 1, 1, 1, time.Local),
					SearchQuery: "zoom.us",
				},
			},
		},
	},
}

func getTime(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *time.Time {
	t := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return &t
}
