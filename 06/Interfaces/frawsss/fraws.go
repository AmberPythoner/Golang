package frawsss

import "fmt"

type INters struct {
	Contens string
}

func (root *INters) String() string {
	return fmt.Sprintf("strings *****", root.Contens)
}

func (root *INters) Get(url string) string {
	return root.Contens
}

func (root *INters) Post(url string, contents map[string]string) {
	root.Contens = contents["names"]
}
