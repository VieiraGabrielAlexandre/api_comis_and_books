package routes

import (
	_ "github.com/VieiraGabrielAlexandre/api_comics_and_books/docs"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/controllers"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/comics", controllers.GetAllPaginate).Methods("GET")
	r.HandleFunc("/comics", controllers.Create).Methods("POST")
	r.HandleFunc("/comics/{id}", controllers.GetByID).Methods("GET")
	r.HandleFunc("/comics/{id}", controllers.Update).Methods("PUT")
	r.HandleFunc("/comics/{id}", controllers.Delete).Methods("DELETE")

	http.Handle("/", r)

	println("Running server on port 8081")
	log.Fatalf(http.ListenAndServe(":8081", nil).Error())
}
