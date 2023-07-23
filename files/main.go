package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func main() {
	total := 0.0
	for _, p := range Products {
		total += p.Price
	}

	// dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:04:05"), total)

	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}

	// err := os.WriteFile("output.txt", []byte(dataStr), 0666)

	// file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	// file, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
	file, err := os.CreateTemp(".", "template-*.json")

	if err == nil {
		defer file.Close()
		// file.WriteString(dataStr)
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error)
	}

	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}

	Printfln("Full path: %v", path)
	Printfln("Volume name: %v", filepath.VolumeName(path))
	Printfln("Dir component: %v", filepath.Dir(path))
	Printfln("File component: %v", filepath.Base(path))
	Printfln("File extension: %v", filepath.Ext(path))
}
