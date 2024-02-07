package routes

import (
	"github.com/gorilla/mux"
	"github.com/joserafaelSH/book_management_system/pkg/controllers"
)

var RegisterBookStoareRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
