package paser

import (
	"crawler/engine"
	"regexp"
)

const cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityReg)
	mathes := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, v := range mathes {
		name := v[2]
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(v[1]),
				ParseFunc: func(bytes []byte) engine.ParseResult {
					return Profile(bytes, string(name))
				},
			})
		result.Item = append(result.Item, "User: "+string(v[2]))

	}
	return result
}
