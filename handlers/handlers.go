package handlers

import (
	"encoding/json"
	"go-rest2/models"
	"log"
	"net/http"
)

type BooksHandler struct {
	store map[string]models.Book
}

func (b *BooksHandler) Get(w http.ResponseWriter, r *http.Request) {
	books := make([]models.Book, len(b.store))
	i := 0
	for _, book := range b.store {
		books[i] = book
		i++
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "[ERROR] error encoding json", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func NewBooksHandler() *BooksHandler {
	return &BooksHandler{
		store: map[string]models.Book{
			"id1": {
				ID:     "1",
				Title:  "Go Programming",
				Author: "John Johnson",
				Year:   2020,
			},
			"id2": {
				ID:     "2",
				Title:  "Go REST",
				Author: "Peter Peterson",
				Year:   2018,
			},
			"id3": {
				ID:     "3",
				Title:  "Golang 'n stuff",
				Author: "Jacob Jacobson",
				Year:   2015,
			},
		},
	}
}
