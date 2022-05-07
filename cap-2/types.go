package main

import "fmt"

var x int = 10

type myType int

var my myType = 10

func main() {
	// static type
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", my)

	// int, string, bool
	// slice, array, struct, map

	x = int(my)
	fmt.Printf("%T", x)

}
