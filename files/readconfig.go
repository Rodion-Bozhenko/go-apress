package main

import (
	"encoding/json"
	"os"
)

type ConfigData struct {
	UserName           string
	AdditionalProducts []Product
}

var Config ConfigData

func LoadConfig() (err error) {
	// data, err := os.ReadFile("config.json")
	file, err := os.Open("config.json")
	if err == nil {
		// decoder := json.NewDecoder(strings.NewReader(string(data)))
		defer file.Close()

		nameSlice := make([]byte, 5)
		file.ReadAt(nameSlice, 17)
		Config.UserName = string(nameSlice)

		file.Seek(49, 0)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config.AdditionalProducts)
	}
	return
}

func init() {
	err := LoadConfig()
	if err != nil {
		Printfln("Error loading config: %v", err.Error())
	} else {
		Printfln("Username: %v", Config.UserName)
		Products = append(Products, Config.AdditionalProducts...)
	}
}
