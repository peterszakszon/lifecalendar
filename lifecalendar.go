package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", calendar)

	fmt.Println("Visit http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func calendar(w http.ResponseWriter, r *http.Request) {
	numYears := 90

	years := make([]int, 0)
	for i := 1; i <= numYears; i++ {
		years = append(years, i)
	}

	weeks := make([]int, 0)
	for i := 1; i <= 52; i++ {
		weeks = append(weeks, i)
	}

	data := new(tdata)
	data.Years = years
	data.Weeks = weeks

	calTemplate.Execute(w, data)
}

var calTemplate = template.Must(template.ParseFiles("lifecalendar.html"))

type tdata struct {
	Years []int
	Weeks []int
}
