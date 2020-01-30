package main

import (
	"crawler/engine"
	"crawler/zhenai/paser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: paser.ParseCityList,
	})
}
