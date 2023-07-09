package main

import "fmt"

func printPrice() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
}

func printPrice2(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "tax:", taxAmount)
}

// Variadic parameter
func printSuppliers(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

// Pointers as parameters
func swapValues(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

// Function's results
func calcTax(price float64) (float64, bool) {
	if price < 100 {
		return price + (price * 0.2), true
	}
	return 0, false
}

func swapValuesUsingMultiResults(first, second int) (int, int) {
	return second, first
}

// Named results
func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func main() {
	fmt.Println("About to call function")
	printPrice()
	fmt.Println("Function complete")

	printPrice2("Kayak", 275, 0.2)
	printPrice2("Lifejacket", 48.95, 0.2)
	printPrice2("Soccer Ball", 19.50, 0.15)

	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers("Lifejacket", "Sail Safe Co")
	printSuppliers("Lifejacket", []string{"Cool Sportgear"}...)
	printSuppliers("Lifejacket")

	// Using pointers as arguments
	val1, val2 := 10, 20
	fmt.Println("Before calling function:", val1, val2)
	// swapValues(&val1, &val2)
	val1, val2 = swapValuesUsingMultiResults(val1, val2)
	fmt.Println("After calling function:", val1, val2)

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		// taxAmount, taxDue := calcTax(price)
		if taxAmount, taxDue := calcTax(price); taxDue {
			fmt.Println("Product:", product, "Tax:", taxAmount)
		} else {
			fmt.Println("Product:", product, "No tax due")
		}
	}

	// Named results

	total1, tax1 := calcTotalPrice(products, 10)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println("Total 2:", total2, "Tax 2:", tax2)
}
