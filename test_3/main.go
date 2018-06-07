package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8085", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles("layout.html")
	// or
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}
