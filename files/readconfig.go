package main

import (
	"encoding/json"
	"os"
	"strings"
)

type ConfigData struct {
	UserName           string
	AdditionalProducts []Product
}

var Config ConfigData

func LoadConfig() (err error) {
	data, err := os.ReadFile("config.json")
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		err = decoder.Decode(&Config)
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
