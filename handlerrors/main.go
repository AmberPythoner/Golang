package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println(reflect.TypeOf(request.URL.Path))
	fmt.Println(request.URL.Path)
	path := request.URL.Path[len("/list/"):]
	open, err := os.Open(path)
	if err != nil {
		return err
	}
	all, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 类似 python中的装饰器
func wrapperHandler(hander appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		error := hander(writer, request)
		if error != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(error):
				code = http.StatusNotFound
			case os.IsPermission(error):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", wrapperHandler(HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
