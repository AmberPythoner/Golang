package paser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
	"strings"
)

const profile = `<div class="des f-cl" data-v-3c42fade>(.*)[元|元以下|元以上]</div>`

func Profile(content []byte, Name string) engine.ParseResult {
	re, _ := regexp.Compile(profile)
	matches := re.FindSubmatch(content)

	result := engine.ParseResult{}
	User := model.Profile{Name: Name}
	//fmt.Printf("%s \n", matches)
	if len(matches) == 0 {
		return engine.ParseResult{}
	}
	values := strings.Split(string(matches[1]), " | ")
	User.Hokou = values[0]
	ages := strings.Trim(values[1], "岁")
	age, _ := strconv.Atoi(ages)
	User.Age = age
	User.Education = values[2]
	User.Marriage = values[3]
	heights := strings.TrimSuffix(values[4], "cm")
	height, _ := strconv.Atoi(heights)
	User.Height = height
	User.Income = values[5]

	result.Item = append(result.Item, User)
	return result
}
