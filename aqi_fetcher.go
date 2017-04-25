package main

import (
	"flag"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/dracher/aqi_fetcher/libs"
)

const (
	// dbPath  = "/home/dracher/GoProjects/src/github.com/dracher/aqi_fetcher/aqidb/aqidata.boltdb"
	feedURL = "http://feed.aqicn.org/feed/%s/en/feed.v1.json"
)

var dbPath = flag.String("d", "/home/dracher/GoProjects/src/github.com/dracher/aqi_fetcher/aqidb/aqidata.boltdb", "the absolute path to boltdb database file")
var cityLists = []string{"beijing", "shanghai", "tianjin", "guangzhou", "chengdu", "hongkong"}

func genFeedURLs() []string {
	feedURLs := []string{}
	for _, city := range cityLists {
		feedURLs = append(feedURLs, fmt.Sprintf(feedURL, city))
	}
	return feedURLs
}

func main() {
	flag.Parse()

	db := libs.NewBoltDB(*dbPath)
	defer db.Close()

	log.Info("Start to fetching aqi data")
	res := libs.FetchAqiData(genFeedURLs())

	log.Warnf("Prepare saving data to %s", *dbPath)
	for _, data := range res {
		libs.SaveAqiData(db, data.Flat())
		log.Infof("Saving %s aqi data to db", data.City.Name)
	}
	log.Warn("All Done~")
}
