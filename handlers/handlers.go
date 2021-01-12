package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest2/models"
	"log"
	"net/http"
	"sync"
	"time"
)

type BooksHandler struct {
	sync.Mutex
	store map[string]models.Book
}

func (b *BooksHandler) ServeBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		b.get(w, r)
		return
	case "POST":
		b.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func (b *BooksHandler) get(w http.ResponseWriter, r *http.Request) {
	books := make([]models.Book, len(b.store))
	b.Lock()
	i := 0
	for _, book := range b.store {
		books[i] = book
		i++
	}
	b.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "[ERROR] error encoding json", http.StatusInternalServerError)
		log.Println(err)
	}
}

func (b *BooksHandler) post(w http.ResponseWriter, r *http.Request) {
	// var body map[string]interface{}
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "[ERROR] error decoding json", http.StatusBadRequest)
		log.Println(err)
		return
	}
	hdr := r.Header.Get("Content-Type")
	if hdr != "application/json" {
		http.Error(w, "[ERROR] bad Content-Type", http.StatusUnsupportedMediaType)
		return
	}
	book.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	b.Lock()
	b.store[book.ID] = book
	defer b.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST success!"))
}

func NewBooksHandler() *BooksHandler {
	return &BooksHandler{
		store: map[string]models.Book{},
	}
}
