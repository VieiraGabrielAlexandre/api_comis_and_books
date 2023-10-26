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
	var page = 1
	var limit = 10

	if r.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	}

	if r.URL.Query().Get("limit") != "" {
		limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	}

	orderBy := r.URL.Query().Get("orderBy")
	filter := r.URL.Query().Get("filter")
	column := r.URL.Query().Get("column")
	order := r.URL.Query().Get("order")

	repo := repositories.ComicsBooksRepositoryDb{Db: database.DB}

	total, _ := repo.Count()
	comics, err := repo.GetAllPaginate(page, limit, filter, orderBy, column, order)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	metadata := struct {
		TotalCount int             `json:"total_count"`
		Page       int             `json:"page"`
		Limit      int             `json:"limit"`
		Comics     []models.Comics `json:"comics"`
	}{
		TotalCount: int(total),
		Page:       page,
		Limit:      limit,
		Comics:     comics,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(metadata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

// Delete handles the deletion of a comic.
func Delete(w http.ResponseWriter, r *http.Request) {
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

	err = repo.Delete(id)

	if err != nil {
		println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Comic updated successfully"))
}
