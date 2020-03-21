package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)


func get_timestamp(t time.Time) string {
    datetime := strings.Split(t.String(), " ")
    date := strings.Split(datetime[0], "-")
    time := strings.Split(datetime[1], ":")

    var YY, MM, DD, HH, mm string
    YY = date[0]
    MM = date[1]
    DD = date[2]
    HH = time[0]
    mm = time[1]
    return YY + MM + DD + "_" + HH + mm
}

func pull_data(input_url, output_file string) {
    var (
        err error
        response *http.Response
        retries int = 3
        timeout time.Duration = 3 * time.Second
    )

    client := http.Client {
        Timeout: time.Duration(timeout),
    }

    // Make request
    for retries > 0 {
        response, err = client.Get(input_url)

        // Fail: Retry
        if err != nil {
            log.Println(err)
            time.Sleep(timeout)
            retries -= 1
        } else {
            // Success: Continue
            break
        }
    }

    if response != nil {
        defer response.Body.Close()

        // Create output file
        outFile, err := os.Create(output_file)
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


type DataDirector struct {
    data_url, out_html string
}


func main() {
    var outdir = "data/"

    now := time.Now()

    data_sites := make([]DataDirector, 3)
    data_sites[0] = DataDirector{"https://www.worldometers.info/coronavirus/",
                                 outdir + "worldometers_" + get_timestamp(now) + ".html"}
    data_sites[1] = DataDirector{"https://www.worldometers.info/coronavirus/country/us/",
                                 outdir + "worldometers_us_" + get_timestamp(now) + ".html"}
    data_sites[2] = DataDirector{"https://ncov2019.live/data",
                                 outdir + "ncov2019_" + get_timestamp(now) + ".html"}


    for _, site := range data_sites {
        pull_data(site.data_url, site.out_html)
    }
}