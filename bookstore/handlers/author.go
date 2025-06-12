package handlers

import (
	"encoding/json"
	"net/http"

	config "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/config"
	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/models"
	repository "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/repository"
)

func GetAuthors( w http.ResponseWriter, r *http.Request )  {
	authorRepo := repository.NewAuthorRepository(config.GetDB())
	authors, err := authorRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(authors)
}


func CreateAuthor( w http.ResponseWriter, r *http.Request )  {
	authorRepo := repository.NewAuthorRepository(config.GetDB())
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := authorRepo.Create(&author); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}