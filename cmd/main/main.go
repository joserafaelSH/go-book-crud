package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	routes "github.com/joserafaelSH/book_management_system/pkg/routes/books"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoareRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
