package postrequests

import (
	"io/ioutil"
	"net/http"
)

type INters struct {
	UserAgent string
}

func (root INters) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
