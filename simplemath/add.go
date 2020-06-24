package main

import (
	"fmt"
)

func Add(num1, num2 int) int {
	var result int
	result = num1 + num2
	return result
}

func main() {
	num1 := 3
	num2 := 5
	result := Add(num1, num2)
	fmt.Printf("result is %d", result)
}
