package main

import (
	"fmt"
	"math"
)

func main() {
	price, tax := 275.00, 27.40

	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64

	fmt.Println("Arifmethic overflow")
	fmt.Println(intVal * 2)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))

	fmt.Println("Reminder operator")

	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))

	fmt.Println(posResult, negResult, absResult)
}
