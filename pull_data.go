package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
    //"github.com/robfig/cron/v3"
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
    data_url string
    out_dir string
    html_fname string
}


func main() {
    var outdir = "data/"

    countries := []string{
        "china",
        "italy",
        "spain",
        "us",
        "germany",
        "iran",
        "france",
        "south-korea",
        "switzerland",
        "uk",
        "netherlands",
        "austria",
        "belgium",
        "norway",
        "sweden",
        "canada",
        "denmark",
        "portugal",
        "malaysia",
        "brazil",
        "australia",
    }

    now := time.Now()
    timestamp := get_timestamp(now)

    data_sites := make([]DataDirector, len(countries))
    for idx, country := range countries {
        data_sites[idx] = DataDirector {
            "https://www.worldometers.info/coronavirus/country/" + country +"/",
            outdir + "worldometers/" + country + "/",
            "worldometers_" + country + "_" + timestamp + ".html",
        }
    }

    data_sites = append(data_sites, DataDirector {
        "https://www.worldometers.info/coronavirus/",
        outdir + "worldometers/global/",
        "worldometers_" + timestamp + ".html",
    })

    data_sites = append(data_sites, DataDirector {
        "https://ncov2019.live/data",
        outdir + "ncov2019/",
        "ncov2019_" + timestamp + ".html",
    })

    for _, site := range data_sites {
        os.MkdirAll(site.out_dir, os.ModePerm)
        pull_data(site.data_url, site.out_dir + site.html_fname)
    }
}