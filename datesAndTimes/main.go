package main

import "time"

func PrintTime(label string, t *time.Time) {
	Printfln("%s: Day: %v Month: %v Year: %v", label, t.Day(), t.Month(), t.Year())
}

func main() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)

	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
}
