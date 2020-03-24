# COVID19-Tracker
Golang Web Scraper &amp; Data Processing COVID19 Project

## Setup
This script requires the cron v3 package, install using the following commands to get past the GOPATH error:
```
$ export GO111MODULE=on

//$ go mod init <project name>
$ go mod init COVID19

//$ go mod download repo@version
$ go mod download github.com/robfig/cron/v3@v3.0.1
```

This script also requires teh goquery package for DOM parsing:
```
$ go get github.com/PuerkitoBio/goquery
```