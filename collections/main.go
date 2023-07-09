package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

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

	// Slices comparison
	equalSlice := allNumbers
	fmt.Println("Equal:", reflect.DeepEqual(allNumbers, equalSlice))

	// Convert slice to sliceToArray
	sliceToArray := []string{"First", "Second", "Third"}
	arrayPtr := (*[3]string)(sliceToArray)
	array := *arrayPtr
	fmt.Println("Array from slice:", array)

	// Maps
	// mapProducts := make(map[string]float64, 10)
	mapProducts := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	mapProducts["Kayak"] = 279
	mapProducts["Lifejacket"] = 48.95
	fmt.Println("Map size:", len(mapProducts))
	fmt.Println("Price:", mapProducts["Kayak"])
	fmt.Println("Price:", mapProducts["Hat"])

	// Removing items
	delete(mapProducts, "Hat")

	if value, ok := mapProducts["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}

	// Enumerating maps
	for key, value := range mapProducts {
		fmt.Println("Product:", key, "Price:", value)
	}

	// Sorting maps

	keys := make([]string, 0, len(mapProducts))
	for key := range mapProducts {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Sorted Product:", key, "Sorted Price:", mapProducts[key])
	}

	// Strings as collections
	var price2 = []rune("£48.95")
	var currency string = string(price2[0])
	var amountString = string(price2[1:])
	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Currency:", currency)

	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}

	// Enumerating strings
	for index, char := range "¢49.95" {
		fmt.Println("Index:", index, "Char:", char, "String char:", string(char))
	}
}
