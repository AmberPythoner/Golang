package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("aa.html")

	//regexRules :=`<div class="m-btn purple" data-v-8b1eac0c>([^\x00-\xff]{2})</div><div class="m-btn purple" data-v-8b1eac0c>([0-9]+)岁</div><div class="m-btn purple" data-v-8b1eac0c>魔羯座(12.22-01.19)</div><div class="m-btn purple" data-v-8b1eac0c>178cm</div><div class="m-btn purple" data-v-8b1eac0c>工作地:阿坝若尔盖</div><div class="m-btn purple" data-v-8b1eac0c>月收入:5-8千</div><div class="m-btn purple" data-v-8b1eac0c>高中及以下</div></div>`
	regexRules := `<div class="des f-cl" data-v-3c42fade>(.*)元</div>`
	//regexRules := `<div class="m-btn purple" data-v-8b1eac0c>([^\x00-\xff]{2,4})</div><div class="m-btn purple" data-v-8b1eac0c>([0-9]+)岁</div><div class="m-btn purple" data-v-8b1eac0c>([^\x00-\xff0-9()]{3,10})</div>`
	re, _ := regexp.Compile(regexRules)
	matches := re.FindSubmatch(bytes)
	fmt.Println(string(matches[1]))
	a := strings.Split(string(matches[1]), "|")
	for _, v := range a {
		fmt.Println(strings.Trim(v, " "))
	}
}
