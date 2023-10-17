package controllers

import "net/http"

func GetAllPaginate(w http.ResponseWriter, r *http.Request) {
	println(w, "GetAllPaginate")
}

func Create(w http.ResponseWriter, r *http.Request) {
	println(w, "Create")
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	println(w, "GetByID")
}

func Update(w http.ResponseWriter, r *http.Request) {
	println(w, "Update")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	println(w, "Delete")
}
