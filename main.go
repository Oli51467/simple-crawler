package main

import (
	"simple-webcrawler/engine"
	"simple-webcrawler/lib/logger"
	"simple-webcrawler/zhenai/parser"
)

const WebsiteUrl = "http://www.zhenai.com/zhenghun"

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "simple-godis",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})
	engine.Run(engine.Request{
		Url:        WebsiteUrl,
		ParserFunc: parser.ParseCityList,
	})
}
