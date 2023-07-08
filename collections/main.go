package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"

	names2 := [3]string{"Alice", "Not Alice", "Another Alice"}

	// Infer array capacity
	names3 := [...]string{"Rodion", "Oleg", "Mayami"}

	// By default Go works with values rather than references
	otherArray := names
	otherPointerArray := &names
	names[0] = "Canoe"

	sameNames := [...]string{"Canoe", "Lifejacket", "Paddle"}
	same := names == sameNames
	notSame := names == otherArray
	fmt.Println(names, otherArray, *otherPointerArray, names2, names3, same, notSame)

	// Enumerating arrays
	for index, name := range names {
		fmt.Println("Index:", index, "Name:", name)
	}

	// Slices
	slice := make([]string, 3, 6)
	slice[0] = "Kayak"
	slice[1] = "Lifejacket"
	slice[2] = "Paddle"

	slice2 := []string{"Slice1", "Slice2", "Slice3"}
	slice2 = append(slice, "Slice4", "LastSlice")

	// slice3 has separate underlying array because with append we get out of slice capacity
	slice3 := append(slice, "Slice4", "slice5", "slice6", "last")

	slice[0] = "Changed Elem"

	fmt.Println("Slice:", slice, slice2, slice3)
	fmt.Println("Len:", len(slice2), "Cap:", cap(slice2))

	anotherSlice := []string{"Some", "Other"}
	slice3 = append(slice3, anotherSlice...)

	fmt.Println(slice3)

	// Making slices from array
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	someNames := products[1:3:4]
	someNames = append(someNames, "another")
	allNames := products[:]

	products[2] = "New"
	fmt.Println(cap(someNames))
	fmt.Println("SomeNames:", someNames, "allNames:", allNames)

	// Copy function
	subProducts := products[1:]
	subSubProducts := make([]string, 2)
	copy(subSubProducts, subProducts)

	fmt.Println("subProducts:", subProducts, "subSubProducts:", subSubProducts)

	allNumbers := []int{12, 42, 55, 69}
	someNumbers := []int{69, 33}
	copy(someNumbers[1:], allNumbers[2:3])
	fmt.Println("allNumbers:", allNumbers, "someNumbers:", someNumbers)

	// Delete element
	sliceForDeletion := []string{"First", "Second", "Third", "Fourth"}
	deleted := append(sliceForDeletion[:2], sliceForDeletion[3:]...)
	fmt.Println("Deleted:", deleted)
}
