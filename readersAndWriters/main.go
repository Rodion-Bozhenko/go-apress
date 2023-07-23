package main

import (
	"bufio"
	"io"
	"strings"
)

type Product struct {
	Name, Category string
	Price          float64
}

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

	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")

	concatReader := io.MultiReader(r1, r2, r3)

	// var w strings.Builder
	// teeReader := io.TeeReader(concatReader, &w)
	//
	// ConsumeData(teeReader)
	// Printfln("Echo data: %v", w.String())

	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)

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

	text := "It was a boat. A small boat."

	var reader io.Reader = NewCustomerReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)

	buffered := bufio.NewReader(reader)

	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}

	Printfln("Read data: %v", writer.String())
}
