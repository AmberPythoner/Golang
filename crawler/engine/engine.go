package engine

import (
	"crawler/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, req := range seeds {
		requests = append(requests, req)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		fmt.Println("Url: ", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("can not get the body %s", r.Url)
			continue
		}
		res := r.ParseFunc(body)
		requests = append(requests, res.Requests...)
		for _, c := range res.Item {
			fmt.Println(c)
		}
	}
}
