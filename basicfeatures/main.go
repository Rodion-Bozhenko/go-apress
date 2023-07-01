package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println("Value:", rand.Int())
	const price, tax float32 = 275.00, 27.50
	const quantity, inStock = 3, true
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("In stock:", inStock)

	var letter = "A"
	fmt.Println("Letter:", letter)
	letter = "B"
	fmt.Println("Letter:", letter)

	short := "short"
	fmt.Println("Short:", short)
	short = "long"
	fmt.Println("Short:", short)

	fmt.Println("Pointers")

	first := 100
	second := first

	first++

	fmt.Println("First:", first, "Second:", second)

	var pointer *int = &first

	first++
	*pointer++
	pointerToPointer := &pointer

	var nilPointer *int

	fmt.Println("First:", first, "Second:", *pointer, "Nil Pointer", nilPointer)
	fmt.Println("Pointer to pointer", **pointerToPointer)

	fmt.Println("Why pointers are useful")

	names := [3]string{"Alice", "Carol", "Bob"}

	secondName := names[1]
	fmt.Println("Second name:", secondName)

	sort.Strings(names[:])
	fmt.Println("Second name after sorting:", secondName)

	secondNames := [3]string{"Alice", "Carol", "Bob"}
	secondNamePointer := &secondNames[1]
	fmt.Println("Second name pointer:", *secondNamePointer)
	sort.Strings(secondNames[:])
	fmt.Println("Second name pointer after sorting:", *secondNamePointer)
}
