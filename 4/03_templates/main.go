package main

import (
	"net/http"
	"html/template"
	"strconv"
)

type Product struct {
	Name string `json:"name"`
	Important bool `json:"important"`
	Done bool `json:"done"`
}

func main() {
	tmpl := template.Must(template.ParseFiles("template.html"))

	products := []Product{
		{"Молоко", false, false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			param := r.FormValue("name")
			index, _ := strconv.ParseInt(param, 10, 0)
			products[index].Done = true
		}
		tmpl.Execute(w, products)
	})

	http.ListenAndServe(":8081", nil)
}

