package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/config/database"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/domains/repositories"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/models"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetAllPaginate(w http.ResponseWriter, r *http.Request) {
	println(w, "GetAllPaginate")
}

func Create(w http.ResponseWriter, r *http.Request) {
	var comic models.Comics

	body, err := io.ReadAll(r.Body)

	if err != nil {
		println(err)
	}

	if err := json.Unmarshal(body, &comic); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%s\n", body)

	repo := repositories.ComicsBooksRepositoryDb{Db: database.DB}

	comic, err = repo.Create(comic)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create !"))
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	println(w, "GetByID")
}

func Update(w http.ResponseWriter, r *http.Request) {
	var comic models.Comics

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) >= 2 {
		idString := parts[len(parts)-1]

		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(r.Body)

		if err != nil {
			println(err)
		}

		if err := json.Unmarshal(body, &comic); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("%s\n", body)

		repo := repositories.ComicsBooksRepositoryDb{Db: database.DB}

		comic, err = repo.Update(comic, id)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Not processed !"))
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Update !"))

		return
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte("Parametro id incorreto !"))

	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	println(w, "Delete")
}
