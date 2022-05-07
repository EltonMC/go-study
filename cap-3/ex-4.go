package main

import "fmt"

type newType int

var newVar newType
var otherVar int

func main() {
	fmt.Printf("%v %T", newVar, newVar)
	newVar = 42
	fmt.Printf("%v %T", newVar, newVar)
	otherVar = int(newVar)
	fmt.Printf("%v %T", otherVar, otherVar)
}
