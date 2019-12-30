package main

import (
	"fmt"
	"interfaces/frawsss"
	"interfaces/postrequests"
)

type Retiver interface {
	Get(string) string
}

func download(r Retiver) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retiver
	r = frawsss.INters{"this is the mock "}
	//fmt.Println(download(r))
	inpter(r)
	fmt.Println("******************************")
	r = postrequests.INters{UserAgent: "Molliza"}
	//fmt.Println(download(r))
	inpter(r)

	// 获取 接口类型的第2种方式
	// type assertion
	if values, ok := r.(postrequests.INters); ok {
		fmt.Println(values.UserAgent)
	} else {
		fmt.Println("no frawsss.INters")
	}
}

func inpter(r Retiver) {
	fmt.Println("%T,  %v", r, r)
	// 获取 接口类型的第1种方式
	switch r.(type) {
	case frawsss.INters:
		fmt.Println("1111")
	case postrequests.INters:
		fmt.Println("22222")
	}
}