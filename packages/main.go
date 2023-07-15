package main

import (
	"fmt"
	_ "packages/data"
	. "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	product := store.Product{
		Name:     "Kayak",
		Category: "Watersports",
	}

	fmt.Println("Name:", product.Name, "Category:", product.Category)

	product2 := store.NewProduct("Kayak", "Watersports", 69.69)
	fmt.Println("Name:", product2.Name, "Category:", product2.Category, "Price:", product2.Price())
	fmt.Println("Price:", ToCurrency(product2.Price()))

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product2, product, *store.NewProduct("Kayak", "Watersports", 420.69)},
	}

	fmt.Println("Name:", cart.CustomerName, "Total:", ToCurrency(cart.GetTotal()))
}
