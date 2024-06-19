package main

import (
	"fmt"

	"github.com/ccsimmi/packages/math"
)


func main() {
	var num1 int = 5
	var num2 int = 4

	addition := math.Add(num1, num2)
	subtraction := math.Subtract(num1, num2)
	multiplication := math.Multiply(num1, num2)

	fmt.Printf("The two numbers are: %d and %d\n", num1, num2)
	fmt.Printf("Addition: %d\nSubtraction: %d\nMultiplication: %d\n", addition, subtraction, multiplication)
}