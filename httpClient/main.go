package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	Printfln("Starting HTTP Server")
	go http.ListenAndServe("localhost:5000", nil)
	time.Sleep(time.Second)

	// response, err := http.Get("http://localhost:5000/html")
	// if err == nil && response.StatusCode == http.StatusOK {
	// 	data, err := io.ReadAll(response.Body)
	// 	if err == nil {
	// 		defer response.Body.Close()
	// 		os.Stdout.Write(data)
	// 	}
	// } else {
	// 	Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	// }

	// response, err := http.Get("http://localhost:5000/json")
	// if err == nil && response.StatusCode == http.StatusOK {
	// 	defer response.Body.Close()
	// 	data := []Product{}
	// 	err = json.NewDecoder(response.Body).Decode(&data)
	// 	if err == nil {
	// 		for _, p := range data {
	// 			Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
	// 		}
	// 	} else {
	// 		Printfln("Decode error: %v", err.Error())
	// 	}
	// } else {
	// 	Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	// }

	// formData := map[string][]string{
	// 	"name":     {"Kayak"},
	// 	"category": {"Watersports"},
	// 	"price":    {"279"},
	// }
	//
	// response, err := http.PostForm("http://localhost:5000/echo", formData)
	//
	// if err == nil && response.StatusCode == http.StatusOK {
	// 	io.Copy(os.Stdout, response.Body)
	// 	defer response.Body.Close()
	// } else {
	// 	Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	// }
	// var builder strings.Builder
	// err := json.NewEncoder(&builder).Encode(Products[0])
	//
	// if err == nil {
	// 	response, err := http.Post("http://localhost:5000/echo", "application/json", strings.NewReader(builder.String()))
	// 	if err == nil && response.StatusCode == http.StatusOK {
	// 		io.Copy(os.Stdout, response.Body)
	// 		defer response.Body.Close()
	// 	} else {
	// 		Printfln("Error: %v", err.Error())
	// 	}
	// } else {
	// 	Printfln("ErrorL %v", err.Error())
	// }

	// var builder strings.Builder
	// err := json.NewEncoder(&builder).Encode(Products[0])
	// if err == nil {
	// 	reqUrl, err := url.Parse("http://localhost:5000/echo")
	// 	if err == nil {
	// 		req := http.Request{
	// 			Method: http.MethodPost,
	// 			URL:    reqUrl,
	// 			Header: map[string][]string{
	// 				"Content-Type": {"application/json"},
	// 			},
	// 			Body: io.NopCloser(strings.NewReader(builder.String())),
	// 		}
	// 		response, err := http.DefaultClient.Do(&req)
	// 		if err == nil && response.StatusCode == http.StatusOK {
	// 			io.Copy(os.Stdout, response.Body)
	// 			defer response.Body.Close()
	// 		} else {
	// 			Printfln("Request Error: %v", err.Error())
	// 		}
	// 	} else {
	// 		Printfln("Parse Error: %v", err.Error())
	// 	}
	// } else {
	// 	Printfln("Encoder Error: %v", err.Error())
	// }

	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/echo", io.NopCloser(strings.NewReader(builder.String())))
	if err == nil {
		req.Header["Content-Type"] = []string{"application/json"}
		response, err := http.DefaultClient.Do(req)
		if err == nil && response.StatusCode == http.StatusOK {
			io.Copy(os.Stdout, response.Body)
			defer response.Body.Close()
		} else {
			Printfln("Request Error: %v", err.Error())
		}
	} else {
		Printfln("Request Init Error: %v", err.Error())
	}
}
