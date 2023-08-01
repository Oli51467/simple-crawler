package parser

import (
	"regexp"
	"simple-webcrawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com.zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList 城市列表解析器
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		// 对于每一个match：
		// 对于每一个城市 将结果追加到ParseResult的Items里
		result.Items = append(result.Items, string(match[2]))
		// 每一个url生成一个request
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: engine.NilParser,
		})
		//logger.Info(fmt.Sprintf("City: %s, Url: %s", match[2], match[1]))
	}
	return result
}
