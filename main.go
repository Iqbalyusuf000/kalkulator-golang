package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", calculatorHandler)
	http.ListenAndServe(":8080", nil)
}

type Calculation struct {
	Num1   float64
	Num2   float64
	Result float64
	Op     string
}

func calculate(num1 float64, num2 float64, op string) float64 {
	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 != 0 {
			return num1 / num2
		}
	}
	return 0
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		num1, _ := strconv.ParseFloat(r.FormValue("num1"), 64)
		num2, _ := strconv.ParseFloat(r.FormValue("num2"), 64)
		op := r.FormValue("operation")
		result := calculate(num1, num2, op)

		data := Calculation{Num1: num1, Num2: num2, Result: result, Op: op}
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, data)
	} else {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	}
}
