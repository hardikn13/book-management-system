package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hardikn13/book-management-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.CollectionOfRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
