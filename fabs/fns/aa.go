package fns

type addrs func() int

func IAddr() addrs {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
