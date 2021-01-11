package main

import (
	"go-rest2/handlers"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	bh := handlers.NewBooksHandler()
	http.HandleFunc("/books", bh.ServeBooks)

	log.Println("Serving on port", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("[ERROR] error starting server.", err)
	}
}
