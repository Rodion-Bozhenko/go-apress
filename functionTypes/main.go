package main

import "fmt"

type calcFunc func(float64) float64

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	// for product, price := range products {
	// 	var calcFunc calcFunc
	// 	fmt.Println("Function assigned:", calcFunc != nil)
	// 	if price > 100 {
	// 		calcFunc = calcWithTax
	// 	} else {
	// 		calcFunc = calcWithoutTax
	// 	}
	// 	fmt.Println("Function assigned:", calcFunc != nil)
	// 	totalPrice := calcFunc(price)
	// 	fmt.Println("Product:", product, "Price:", totalPrice)
	// }
	// for product, price := range products {
	// 	if price > 100 {
	// 		printPrice(product, price, calcWithTax)
	// 	} else {
	// 		printPrice(product, price, calcWithoutTax)
	// 	}
	// }

	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

// Functions as arguments
func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

// Functions as a result
//
//	func selectCalculator(price float64) calcFunc {
//		if price > 100 {
//			return calcWithTax
//		}
//		return calcWithoutTax
//	}

// Literal function syntax
func selectCalculator(price float64) calcFunc {
	if price > 100 {
		// var withTax calcFunc = func(price float64) float64 {
		// 	return price + (price * 0.2)
		// }
		// return withTax
		return func(price float64) float64 {
			return price + (price * 0.2)
		}
	}
	// withoutTax := func(price float64) float64 {
	// 	return price
	// }
	// return withoutTax
	return func(price float64) float64 {
		return price
	}

}
