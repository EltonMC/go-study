package main

import "fmt"

var y = "hello"

func main() {
	x := 10.2

	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("y: %v, %T \n", y, y)

	x = 20

	fmt.Printf("x: %v, %T \n", x, x)
}
