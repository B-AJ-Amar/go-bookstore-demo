package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	config "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/config"
	models "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/models"
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

func GetBook( w http.ResponseWriter, r *http.Request )  {
	bookRepo := repositories.NewBookRepository(config.GetDB())
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	book, err := bookRepo.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func UpdateBook( w http.ResponseWriter, r *http.Request )  {
	bookRepo := repositories.NewBookRepository(config.GetDB())
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	book.ID = uint(id)
	if err := bookRepo.Update(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func DeleteBook( w http.ResponseWriter, r *http.Request )  {
	bookRepo := repositories.NewBookRepository(config.GetDB())
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	if err := bookRepo.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}