package paser

import (
	"crawler/engine"
	"regexp"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-5e16505f>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListReg)
	mathes := re.FindAllSubmatch(content, -1)
	var limit = 6
	result := engine.ParseResult{}
	for _, v := range mathes {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       string(v[1]),
				ParseFunc: ParseCity,
			})
		result.Item = append(result.Item, "City: "+string(v[2]))
		limit--
		if limit < 0 {
			break
		}
	}
	return result
}
