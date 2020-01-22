package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

const prefix = "/list/"

type UserError string

func (erre UserError) Message() string {
	return string(erre)
}

func (erre UserError) Error() string {
	return erre.Message()
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println(reflect.TypeOf(request.URL.Path))
	fmt.Println(request.URL.Path)
	if strings.Index(request.URL.Path, prefix) != 0 {
		fmt.Println("the path must start whit" + prefix)
		return UserError(fmt.Sprintf("the path must start with %s", prefix))
	}
	path := request.URL.Path[len(prefix):]
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

type userError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 类似 python中的装饰器
func wrapperHandler(hander appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("aaa")
				if errormes, ok := r.(userError); ok {
					http.Error(writer, errormes.Message(), http.StatusBadGateway)
				}

			}
		}()

		error := hander(writer, request)
		fmt.Println("this is in errors", error)
		// 如果没有 recover 错误会一直上报 直到程序停止
		if error != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(error):
				code = http.StatusNotFound
			case os.IsPermission(error):
				code = http.StatusForbidden
				//default:
				//	code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", wrapperHandler(HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
