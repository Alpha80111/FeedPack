package main

import (
	"enterpret/config"
	"enterpret/dataaccess"
	"enterpret/server"
	sources2 "enterpret/sources"
	"log"
	"sync"
)

func main() {
	ds := dataaccess.NewDataStore()
	sourceProcessor := sources2.NewSourceProcessor(ds)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		err := server.NewServer(ds)
		if err != nil {
			log.Default().Println("Server ended with error ", err.Error())
		}
		wg.Done()
	}()

	for _, t := range config.TC {
		for _, s := range t.Sources {
			fp, err := sourceProcessor.GetProcessor(s.Name)
			if err != nil {
				log.Default().Println("Failed fetching message processor for source "+s.Name, err.Error())
				return
			}

			_, err = fp.FetchAndStoreFeedbacks(s.Params, t.Name)
			if err != nil {
				log.Default().Println("Failed fetching and storing messages: ", err.Error())
				return
			}
		}
	}

	wg.Wait()
}
