package main

func main() {
	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: %v", p.Name, p.Category, p.Price)
	}
}
