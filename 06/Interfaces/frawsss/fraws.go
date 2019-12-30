package frawsss

type INters struct {
	Contens string
}

func (root INters) Get(url string) string {
	return root.Contens
}
