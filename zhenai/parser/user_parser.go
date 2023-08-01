package parser

import (
	"regexp"
	"simple-webcrawler/engine"
	"simple-webcrawler/model"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<+])</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var nativeSpaceRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)

func ParseUser(contents []byte, name string) engine.ParseResult {
	user := model.User{}
	user.Name = name
	user.Age = extractInt(contents, ageRe)
	user.Height = extractInt(contents, heightRe)
	user.Weight = extractInt(contents, weightRe)

	user.Marriage = extractString(contents, marriageRe)
	user.Income = extractString(contents, incomeRe)
	user.Gender = extractString(contents, genderRe)
	user.Car = extractString(contents, carRe)
	user.House = extractString(contents, houseRe)
	user.Education = extractString(contents, educationRe)
	user.Constellation = extractString(contents, constellationRe)
	user.NativeSpace = extractString(contents, nativeSpaceRe)
	user.Occupation = extractString(contents, occupationRe)

	result := engine.ParseResult{
		Items: []interface{}{user},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		res, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return res
		}
		return 0
	} else {
		return 0
	}
}
