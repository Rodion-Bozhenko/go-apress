package main

import (
	"fmt"
	"strconv"
)

func main() {
	price := 275.00

	if price > 500 {
		fmt.Println("Price is large indeed")
	} else if price < 300 {
		fmt.Println("Price is low indeed")
	}

	priceString := "254"

	if price, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", price)
	} else {
		fmt.Println("Error:", err)
	}

	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}

	counter2 := 0
	for counter2 <= 3 {
		fmt.Println("Counter2:", counter2)
		counter2++
	}

	for counter := 0; counter <= 3; counter++ {
		fmt.Println("Counter3:", counter)
	}

	for counter := 0; counter <= 3; counter++ {
		if counter == 1 {
			continue
		}
		fmt.Println("Counter4:", counter)
	}

	product := "Kayak"

	for index, character := range product {
		fmt.Println("Index:", index, "Character:", string(character))
	}

	products := [5]string{"Kayak", "Lifejacket", "Soccer Ball"}

	products[3] = "kayak"
	products[4] = "lifejacket"
	for _, product := range products {
		fmt.Println("Product:", product)
	}

	for index, p := range products {
		switch p {
		case "Kayak", "kayak":
			if p == "kayak" {
				fmt.Println("kayak:", index)
				break
			}
			fmt.Println("Kayak:", index)
		case "Lifejacket":
			fmt.Println("Lifejacket:", index)
			fallthrough
		case "lifejacket":
			fmt.Println("lifejacket:", index)
		default:
			fmt.Println("Product:", p, "at index:", index)
		}
	}

	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime number:", val)
		default:
			fmt.Println("Non-prime number:", val)
		}
	}

	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero")
		case counter < 3:
			fmt.Println("is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println("between 2 and 7")
		default:
			fmt.Println("Default: ", counter)
		}
	}

	targetCounter := 0
target:
	fmt.Println("Target counter:", targetCounter)
	targetCounter++
	if targetCounter < 5 {
		goto target
	}
}
