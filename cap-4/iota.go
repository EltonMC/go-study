package main

import "fmt"

const (
	a = iota + 1000
	b
	c
	_
	f
)

func main() {
	fmt.Println(a, b, c, f)

	const g = iota

	fmt.Println(g)
}

