package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

func IntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}

func main() {
	val1 := 279.00
	val2 := 70.5

	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("Copysign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, 2))
	Printfln("Round: %v", math.Round(val2))
	Printfln("RoundToEven: %v", math.RoundToEven(val2))

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		// Printfln("Value %v : %v", i, rand.Int())
		// Printfln("Value %v : %v", i, rand.Intn(10))
		Printfln("Value %v : %v", i, IntRange(10, 20))
	}

	rand.Shuffle(len(names), func(first, second int) {
		names[first], names[second] = names[second], names[first]
	})

	for i, name := range names {
		Printfln("Index %v: Name: %v", i, name)
	}

	ints := []int{9, 4, 4, 2, 6, 9, 0}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints sorted: %v", ints)

	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats sorted: %v", floats)

	strings := []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	sort.Strings(strings)
	Printfln("Strings sorted: %v", strings)
}
