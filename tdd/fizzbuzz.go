package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Insira um número inteiro: ")
	var input int
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(i int) string {
	var result string
	if i%3 == 0 {
		result = "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		result = strconv.Itoa(i)
	}
	return result
}
