package main

import "fmt"

type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}

	return totals
}

type Supplier struct {
	name, city string
}

func printDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.price)
}

func (product *Product) printDetailsMethod() {
	fmt.Println("Name2:", product.name, "Category2:", product.category, "Price2:", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshhold float64) float64 {
	if product.price > threshhold {
		return product.price + (product.price * rate)
	}

	return product.price
}

func (supplier *Supplier) printDetailsMethod() {
	fmt.Println("Name:", supplier.name, "City:", supplier.city)
}

func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 69.420},
		{"Soccer Ball", "Soccer", 19.50},
	}
}

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}

	return
}

type Account struct {
	accountNumber int
	expenses      []Expense
}

type Person struct {
	name, city string
}

func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name)
	case *Product:
		fmt.Println("Product pointer:", value.name)
	case Service:
		fmt.Println("Service:", value.description, "Price:",
			value.monthlyFee*float64(value.durationMonths))
	case Person:
		fmt.Println("Person:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Person Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default:", value)
	}
}

func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for _, p := range products {
		printDetails(p)
		p.printDetailsMethod()
	}

	suppliers := []*Supplier{
		{"Acme Co", "New York"},
		{"Boat Co", "Chicago"},
	}

	for _, s := range suppliers {
		s.printDetailsMethod()
	}

	bat := Product{"Bat", "Sport", 69}
	bat.printDetailsMethod()

	products2 := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 69.420},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for category, total := range products2.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}

	products3 := ProductList(getProducts())

	for category, total := range products3.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}

	product := Product{"Kayak", "Watersports", 420}
	insurance := Service{"Boat Cover", 12, 69.69}

	fmt.Println("Product:", product.name, "Service:", insurance.description, "Price:", insurance.monthlyFee*float64(insurance.durationMonths))

	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 69.420},
	}

	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

	fmt.Println("Total:", calcTotal(expenses))

	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 69.420},
		},
	}

	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

	fmt.Println("Total:", calcTotal(account.expenses))

	// Can use pointers to the struct when assigning to interface variable
	var expense Expense = &product

	product.price = 69

	fmt.Println("Product price:", product.price, "Expense price:", expense.getCost(false))

	var e1 Expense = &Product{name: "Kayak"}
	var e2 Expense = &Product{name: "Kayak"}

	var e3 Expense = Service{description: "Boat Cover"}
	var e4 Expense = Service{description: "Boat Cover"}

	fmt.Println("e1 == e2:", e1 == e2)
	fmt.Println("e3 == e4:", e3 == e4)

	// type assertion

	expenses2 := []Expense{
		Service{"Boat Cover", 69, 420.69},
		Service{"Paddle Protect", 42, 69},
		&Product{"Kayak", "Watersports", 123},
	}

	for _, expense := range expenses2 {
		// if s, ok := expense.(Service); ok {
		// 	fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
		// } else {
		// 	fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		// }

		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description)
		case *Product:
			fmt.Println("Product:", value.name)
		default:
			fmt.Println("Expense:", value.getName())
		}
	}

	var expense2 Expense = &Product{"Kayak", "Watersports", 275}

	data := []interface{}{
		expense2,
		Product{"Lifejacket", "Watersports", 69},
		Service{"Boat Cover", 12, 89.50},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is string",
		100,
		true,
	}

	for _, item := range data {
		processItem(item)
	}
}
