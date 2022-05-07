package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

func main() {
	s := fmt.Sprintf("%v %v %v", a, b, c)
	fmt.Println(s)
}
