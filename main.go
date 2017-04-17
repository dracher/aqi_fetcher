package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/dracher/aqi_fetcher/libs"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var cityLists = []string{
	"http://feed.aqicn.org/feed/beijing/en/feed.v1.json",
	"http://feed.aqicn.org/feed/shanghai/en/feed.v1.json",
	"http://feed.aqicn.org/feed/tianjin/en/feed.v1.json",
	"http://feed.aqicn.org/feed/guangzhou/en/feed.v1.json",
	"http://feed.aqicn.org/feed/chengdu/en/feed.v1.json",
	"http://feed.aqicn.org/feed/hongkong/en/feed.v1.json"}

var engine *xorm.Engine

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./aqidata.sqlite")
	if err != nil {
		log.Panic(err)
	}
	for _, city := range []string{"beijing", "tianjin", "shanghai", "guangzhou", "chengdu", "hongkong"} {
		engine.Sync(libs.TableData{CityName: city})
	}

	res := libs.FetchAqiData(cityLists)
	libs.SaveToDB(res, engine)
}
