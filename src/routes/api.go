package routes

import (
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/comics", controllers.GetAllPaginate).Methods("GET")
	r.HandleFunc("/comics", controllers.Create).Methods("POST")
	r.HandleFunc("/comics/{id}", controllers.GetByID).Methods("GET")
	r.HandleFunc("/comics/{id}", controllers.Update).Methods("PUT")
	r.HandleFunc("/comics/{id}", controllers.Delete).Methods("DELETE")

	http.Handle("/", r)

	println("Running server on port 8081")
	log.Fatalf(http.ListenAndServe(":8081", nil).Error())
}
