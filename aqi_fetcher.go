package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/dracher/aqi_fetcher/libs"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath  = "/home/dracher/GoProjects/src/github.com/dracher/aqi_fetcher/aqidata.sqlite"
	feedURL = "http://feed.aqicn.org/feed/%s/en/feed.v1.json"
)

var cityLists = []string{"beijing", "shanghai", "tianjin", "guangzhou", "chengdu", "hongkong"}

func genFeedURLs() []string {
	feedURLs := []string{}
	for _, city := range cityLists {
		feedURLs = append(feedURLs, fmt.Sprintf(feedURL, city))
	}
	return feedURLs
}

func initDB() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", dbPath)
	if err != nil {
		log.Panic(err)
	}
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Warnf("%s not exists, start to create table schema", dbPath)
		for _, city := range cityLists {
			engine.Sync(libs.TableData{CityName: city})
		}
	}
	return engine
}

func main() {
	orm := initDB()
	res := libs.FetchAqiData(genFeedURLs())
	libs.SaveToDB(res, orm)
}
