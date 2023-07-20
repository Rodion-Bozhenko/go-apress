package main

import "fmt"

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
	} else {
		err = fmt.Errorf("Error for index %v", index)
	}
	return
}

func main() {
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")

	fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)

	productName, err := getProductName(10)

	if err == nil {
		fmt.Println(productName)
	} else {
		fmt.Println(err.Error())
	}

	number := 250

	fmt.Printf("Binary: %b", number)
	fmt.Printf("Decimal: %d", number)
	fmt.Printf("Octal: %o %O", number, number)
	fmt.Printf("Hexadecimal: %x, %X", number, number)
}
