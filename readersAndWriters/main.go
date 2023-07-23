package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Product struct {
	Name, Category string
	Price          float64
}

type DiscountProduct struct {
	*Product `json:"product,omitempty"`
	Discount float64 `json:",string"`
}

func (dp *DiscountProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}

type Named interface{ GetName() string }

type Person struct{ PersonName string }

func (p *Person) GetName() string { return p.PersonName }

func (p *DiscountProduct) GetName() string { return p.Name }

var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    279}

var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func copyData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	// var builder strings.Builder
	// r := strings.NewReader("Kayak")
	// processData(r, &builder)
	// Printfln("String builder contents: %s", builder.String())
	// copyData(r, &builder)
	// Printfln("String builder contents from copy: %s", builder.String())

	// pipeReader, pipeWriter := io.Pipe()
	// go GenerateData(pipeWriter)
	// ConsumeData(pipeReader)

	// r1 := strings.NewReader("Kayak")
	// r2 := strings.NewReader("Lifejacket")
	// r3 := strings.NewReader("Canoe")
	//
	// concatReader := io.MultiReader(r1, r2, r3)

	// var w strings.Builder
	// teeReader := io.TeeReader(concatReader, &w)
	//
	// ConsumeData(teeReader)
	// Printfln("Echo data: %v", w.String())

	// limited := io.LimitReader(concatReader, 5)
	// ConsumeData(limited)

	// var w1 strings.Builder
	// var w2 strings.Builder
	// var w3 strings.Builder
	//
	// combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	//
	// GenerateData(combinedWriter)
	//
	// Printfln("Writer #1: %v", w1.String())
	// Printfln("Writer #2: %v", w1.String())
	// Printfln("Writer #3: %v", w1.String())

	// text := "It was a boat. A small boat."

	// var reader io.Reader = NewCustomerReader(strings.NewReader(text))
	// var writer strings.Builder
	// slice := make([]byte, 5)
	//
	// buffered := bufio.NewReader(reader)
	//
	// for {
	// 	count, err := buffered.Read(slice)
	// 	if count > 0 {
	// 		Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
	// 		writer.Write(slice[0:count])
	// 	}
	// 	if err != nil {
	// 		break
	// 	}
	// }
	//
	// Printfln("Read data: %v", writer.String())

	// var builder strings.Builder
	// var writer = bufio.NewWriterSize(NewCustomerWriter(&builder), 20)
	// for i := 0; true; {
	// 	end := i + 5
	// 	if end >= len(text) {
	// 		writer.Write([]byte(text[i:]))
	// 		writer.Flush()
	// 		break
	// 	}
	// 	writer.Write([]byte(text[i:end]))
	// 	i = end
	// }
	// Printfln("Written data: %v", builder.String())

	var b bool = true
	var str string = "Hello"
	var fval float64 = 69.420
	var ival int = 200
	var pointer *int = &ival
	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}
	dp := DiscountProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}
	dp2 := DiscountProduct{
		Discount: 10.3,
	}

	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}

	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)
	encoder.Encode(m)
	encoder.Encode(Kayak)
	encoder.Encode(dp)
	encoder.Encode(dp2)

	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)

	fmt.Print(writer.String())

	reader := strings.NewReader(`true "Hello" 99.99 20`)

	vals := []interface{}{}

	decoder := json.NewDecoder(reader)

	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}

	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}
}
