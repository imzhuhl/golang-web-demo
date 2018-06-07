package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/", allBook).Methods("GET")
	bookRouter.HandleFunc("/{title}/page/{page}", oneBook)

	// r is handler
	http.ListenAndServe(":8085", r)
}

func allBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all book page")
}

func oneBook(w http.ResponseWriter, r *http.Request) {
	// we get param from Vars
	vars := mux.Vars(r)
	fmt.Fprintf(w, "title: %s\tpage: %s", vars["title"], vars["page"])
}
