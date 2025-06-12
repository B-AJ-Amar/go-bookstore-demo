package handlers

import (
	"encoding/json"
	"net/http"

	config "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/config"
	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/models"
	repositories "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/repositories"
)

func GetBooks( w http.ResponseWriter, r *http.Request )  {
	bookRepo := repositories.NewBookRepository(config.GetDB())
	books, err := bookRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}


func CreateBook( w http.ResponseWriter, r *http.Request )  {
	bookRepo := repositories.NewBookRepository(config.GetDB())
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := bookRepo.Create(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}