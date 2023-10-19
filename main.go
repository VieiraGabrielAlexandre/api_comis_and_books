package main

import (
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/config/database"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/config/environments"
	"github.com/VieiraGabrielAlexandre/api_comics_and_books/src/routes"
)

func main() {
	environments.Set()
	database.Connect()
	routes.HandleRequests()
}
