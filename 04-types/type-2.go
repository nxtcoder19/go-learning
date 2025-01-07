package main

import "fmt"

type domain interface {
	square() int
}

type Impl struct {
	a, b int
}

func (i *Impl) square() int {
	return i.a * i.b
}

func newImpl(a, b int) domain {
	return &Impl{a, b}
}

func main() {
	d := newImpl(3, 4)
	fmt.Println(d.square())
}
