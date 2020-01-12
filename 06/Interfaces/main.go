package main

import (
	"fmt"
	"interfaces/frawsss"
)

type Retiver interface {
	Get(string) string
}

type Postvier interface {
	Post(url string, contents map[string]string)
}

type PostAndGet interface {
	Retiver
	Postvier
}

const url = "https://www.imooc.com"

func download(r Retiver) string {
	return r.Get(url)
}

func sessions(r PostAndGet) {
	r.Post(url, map[string]string{
		"names": "ccmuse",
	})
	r.Get(url)
	fmt.Println(r)
}

func main() {
	var r PostAndGet
	r = &frawsss.INters{"this is the mock "}
	sessions(r)

	////fmt.Println(download(r))
	//inpter(r)
	//fmt.Println("******************************")
	//r = postrequests.INters{UserAgent: "Molliza"}
	////fmt.Println(download(r))
	//inpter(r)
	//
	//// 获取 接口类型的第2种方式
	//// type assertion
	//if values, ok := r.(postrequests.INters); ok {
	//	fmt.Println(values.UserAgent)
	//} else {
	//	fmt.Println("no frawsss.INters")
	//}
}

//
//func inpter(r Retiver) {
//	fmt.Println("%T,  %v", r, r)
//	// 获取 接口类型的第1种方式
//	switch r.(type) {
//	case frawsss.INters:
//		fmt.Println("1111")
//	case postrequests.INters:
//		fmt.Println("22222")
//	}
//}
