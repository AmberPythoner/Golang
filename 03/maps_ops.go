package main

import (
	"fmt"
)

func subStrings(s string) {
	start, maxLength := 0, 0
	var lastcood = make(map[byte]int)
	for i, ch := range []byte(s) {
		//fmt.Println(reflect.TypeOf(i))
		//fmt.Println(reflect.TypeOf(ch))
		if v, ok := lastcood[ch]; ok && v >= start {
			start = i
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastcood[ch] = i
	}
	fmt.Println(maxLength)
}

func main() {
	subStrings("helwo")
}
