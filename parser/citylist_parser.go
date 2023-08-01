package parser

import (
	"fmt"
	"regexp"
	"simple-webcrawler/lib/logger"
)

// ParseCityList 城市列表解析器
func ParseCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com.zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		logger.Info(fmt.Sprintf("City: %s, Url: %s", match[2], match[1]))
	}
}
