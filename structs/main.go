package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func main() {
	// kayak := Product{
	// 	name:     "Kayak",
	// 	category: "Watersports",
	// }

	// Omitting values names

	acme := &Supplier{"Acme Co", "New York"}

	kayak := Product{"Kayak", "Watersports", 275.00, acme}

	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)

	var lifejacket Product
	fmt.Println("Name is zero:", lifejacket.name == "")
	fmt.Println("Category is zero:", lifejacket.category == "")
	fmt.Println("Price is zero:", lifejacket.price == 0)

	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}

	stockItem := StockLevel{
		Product:   Product{"Kayak", "Watersports", 275.00, acme},
		Alternate: Product{"Alt", "Alt Category", 12, acme},
		count:     100,
	}

	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Name:", stockItem.Alternate.name)
	fmt.Println("Count:", stockItem.count)

	type Item struct {
		name, category string
		price          float64
		*Supplier
	}

	prod := Product{"Kayak", "Watersports", 275.00, acme}
	item := Item{"Kayak", "Watersports", 275.00, acme}

	fmt.Println("Is struct converted:", prod == Product(item))

	// Anonymous struct
	writeName(prod)
	writeName(item)

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})

	fmt.Println(builder.String())

	array := [1]StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00, acme},
			Alternate: Product{"Lifejacket", "Watersports", 48.95, acme},
			count:     100,
		},
	}

	fmt.Println("Array:", array[0].Product.name)

	slice := []StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00, acme},
			Alternate: Product{"Lifejacket", "Watersports", 48.95, acme},
			count:     100,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)

	kvp := map[string]StockLevel{
		"kayak": {
			Product:   Product{"Kayak", "Watersports", 275.00, acme},
			Alternate: Product{"Lifejacket", "Watersports", 48.95, acme},
			count:     100,
		},
	}
	fmt.Println("Map:", kvp["kayak"].Product.name)

	// Pointers

	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275.00,
	}

	p2 := p1
	p3 := &p1

	p1.name = "Original Kayak"

	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", p2.name)
	fmt.Println("P3:", (*p3).name)

	kayak2 := calcTax(&Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	})

	fmt.Println("Name:", kayak2.name, "Category:", kayak2.category, "Price:", kayak2.price)

	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275, acme),
		newProduct("Hat", "Skiing", 42.50, acme),
	}

	for _, product := range products {
		fmt.Println("Name:", product.name, "Supplier:", product.Supplier.name, product.Supplier.city)
	}

	product1 := newProduct("Kayak", "Watersports", 275, acme)
	product2 := deepCopyProduct(product1)

	product1.name = "Original Kayak"
	product1.Supplier.name = "Boat Co"

	for _, product := range []Product{*product1, product2} {
		fmt.Println("Name:", product.name, "Supplier:", product.Supplier.name, product.Supplier.city)
	}
}

func writeName(val struct {
	name, category string
	price          float64
	*Supplier
}) {
	fmt.Println("Anonymous struct name:", val.name)
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

// Struct Constructor Function
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price, supplier}
}

func deepCopyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}
