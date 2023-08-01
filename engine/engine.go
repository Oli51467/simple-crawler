package engine

import (
	"fmt"
	"simple-webcrawler/fetcher"
	"simple-webcrawler/lib/logger"
)

// Run 爬虫引擎 维护一个拉取队列 将每一个要拉取的网页送给拉取器和解析器
func Run(seeds ...Request) {
	var requestQueue []Request
	for _, request := range seeds {
		requestQueue = append(requestQueue, request)
	}

	for len(requestQueue) > 0 {
		request := requestQueue[0]
		requestQueue = requestQueue[1:]

		logger.Info("fetching: " + request.Url)
		fetchRawResult, err := fetcher.Fetch(request.Url)
		if err != nil {
			logger.Error(fmt.Sprintf("Fetch error "+"fetching url %s %v", request.Url, err))
			continue
		}
		// 通过特定的解析函数去解析
		parseResult := request.ParserFunc(fetchRawResult)
		// 解析出的结果如果存在新的url，继续加入到引擎队列中
		requestQueue = append(requestQueue, parseResult.Requests...)
		// 打印Items
		for _, item := range parseResult.Items {
			logger.Info(fmt.Sprintf("Items: %s, ", item))
		}

	}
}
