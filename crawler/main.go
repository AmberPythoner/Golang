package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/paser"
)

func main() {
	e := engine.ConcorrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkInt:   10,
	}
	e.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: paser.ParseCityList,
	})
}
