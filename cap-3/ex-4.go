package main

import "fmt"

type newType int

var newVar newType

func main() {
	fmt.Printf("%v %T", newVar, newVar)
	newVar = 42
	fmt.Printf("%v %T", newVar, newVar)

}
