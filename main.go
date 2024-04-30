package main

import (
	"net/http"

	"github.com/DuarteDc/rest-go/db"
	"github.com/DuarteDc/rest-go/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.Connection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
