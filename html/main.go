package main

import (
	"html/template"
	"os"
)

func main() {
	// templates, err := template.ParseFiles("templates/template.html", "templates/extras.html")
	templates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		templates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		templates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
		for _, t := range templates.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}
}
