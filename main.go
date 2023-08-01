package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"simple-webcrawler/lib/logger"
	"simple-webcrawler/parser"
)

const WebsiteUrl = "http://www.zhenai.com/zhenghun"

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "simple-godis",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})
	resp, err := http.Get(WebsiteUrl)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		logger.Error("Error: status code: ", resp.StatusCode)
		return
	} else {
		rawData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		parser.ParseCityList(rawData)
	}
}
