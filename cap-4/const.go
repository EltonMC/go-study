package main

import "fmt"

// o tipo é atribuído no uso
const x = 10

// o tipo é atribuído aqui
var y = 10

var d float64

func main() {
	d = x
	fmt.Println(d)
}
