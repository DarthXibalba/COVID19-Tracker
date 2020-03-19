package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type DataDirector struct {
	data_url, out_html string
}

func main() {
	var out_dir = "data/"

	var date, time string
	date = "20200319"
	time = "1626"
	datetime := date + "_" + time

	data_sites := make([]DataDirector, 2)
	data_sites[0] = DataDirector{"https://www.worldometers.info/coronavirus/",
								out_dir + "worldometers_" + datetime + ".html"}
	data_sites[1] = DataDirector{"https://ncov2019.live/data", 
								out_dir + "ncov2019_" + datetime + ".html"}

	for _, site := range data_sites {
		// Make request
		response, err := http.Get(site.data_url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		// Create output file
		outFile, err := os.Create(site.out_html)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()

		// Copy data from HTTP response to file
		_, err = io.Copy(outFile, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}