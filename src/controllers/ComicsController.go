package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/config/database"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/domains/repositories"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/models"
)

// Create handles the creation of a comic.
func Create(w http.ResponseWriter, r *http.Request) {
	var comic models.Comics

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &comic); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%s\n", body)

	repo := repositories.ComicsBooksRepositoryDb{Db: database.DB}
	comic, err = repo.Create(comic)

	if err != nil {
		http.Error(w, "Failed to create comic", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Comic created successfully"))
}

// GetByID retrieves a comic by its ID.
func GetByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 2 {
		http.Error(w, "Incorrect ID parameter", http.StatusUnprocessableEntity)
		return
	}

	idString := parts[len(parts)-1]
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	repo := repositories.ComicsBooksRepositoryDb{Db: database.DB}
	comic, err := repo.GetByID(id)

	if err != nil {
		http.Error(w, "Failed to retrieve comic", http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(comic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Update handles the updating of a comic.
func Update(w http.ResponseWriter, r *http.Request) {
	var comic models.Comics

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 2 {
		http.Error(w, "Incorrect ID parameter", http.StatusUnprocessableEntity)
		return
	}

	idString := parts[len(parts)-1]
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &comic); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%s\n", body)

	repo := repositories.ComicsBooksRepositoryDb{Db: database.DB}
	comic, err = repo.Update(comic, id)

	if err != nil {
		http.Error(w, "Failed to update comic", http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Comic updated successfully"))
}

// GetAllPaginate retrieves comics with pagination.
func GetAllPaginate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetAllPaginate")
}

// Delete handles the deletion of a comic.
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete")
}
