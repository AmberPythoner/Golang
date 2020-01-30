package main

import (
	"io"
	"net/http"
	_ "net/http/pprof"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":8888", nil)
}
