package parser

import (
	"regexp"
	"simple-webcrawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^<]*>([^<]+)</a>`

// ParseCity 对每个CityUrl解析其中的用户信息
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		name := string(match[2])
		// 对于每一个match：
		// 对于每一个用户 将结果追加到ParseResult的Items里
		result.Items = append(result.Items, "User "+name)
		// 每一个url生成一个request 这个request是个用户的url
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseUser(bytes, name)
			},
		})
		//logger.Info(fmt.Sprintf("City: %s, Url: %s", match[2], match[1]))
	}
	return result
}
