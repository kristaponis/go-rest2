package main

import (
	"go-rest2/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/books", handlers.BooksHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("[ERROR] error starting server.", err)
	}
}
