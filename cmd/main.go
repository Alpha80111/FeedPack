package main

import (
	"enterpret/config"
	"enterpret/dataaccess"
	"enterpret/models"
	"enterpret/server"
	"enterpret/sources/discourse"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	sources := []config.SourceConfig{
		{
			Name: "discourse",
		},
		{
			Name: "template",
		},
	}

	_ = sources

	ds := dataaccess.NewDataStore()
	dMp := discourse.NewDiscourseMessageProcessor(ds)

	_, err := dMp.FetchAndStoreMessages(models.Params{
		Since:       nil,
		Before:      nil,
		SearchQuery: "zoom.us",
		Source:      "discourse",
	}, "zoom")

	if err != nil {
		fmt.Println("Failed fetching and storing messages: ", err.Error())
		return
	}

	fmt.Println(ds.FetchMessages(dataaccess.GetSourceOption("discourse"), dataaccess.GetTenantOption("zoom")))

	server.NewServer(ds)

	select {}
}
