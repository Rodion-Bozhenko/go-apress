package main

import (
	"composition/store"
	"fmt"
)

func main() {
	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}

	for _, p := range []*store.Product{kayak, lifejacket} {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}

	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 69, 2, true),
	}

	// Fields promotion
	for _, b := range boats {
		fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name, "Price:", b.Price(0.2))
	}

	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}

	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2), "Captain:", r.Captain)
	}

	deal := store.NewSpecialDeal("Weekend Special", kayak, 50)

	Name, price, Price := deal.GetDetails()
	fmt.Println("Name:", Name, "Price field:", price, "Price method:", Price)

	// Promotion ambiguity
	// type OfferBundle struct {
	// 	*store.SpecialDeal
	// 	*store.Product
	// }
	//
	// bundle := OfferBundle{
	// 	store.NewSpecialDeal("Weekend Special", kayak, 50),
	// 	store.NewProduct("Lifejacket", "Watersports", 48.95),
	// }
	//
	// fmt.Println("Price:", bundle.Price(0))

	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}

	for _, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("Name Describable:", item.GetName(), "Category:", item.GetCategory(), "Price:", item.Price(0.2))
		}
	}

}
