package routes

import (
	"github.com/gorilla/mux"
	"github.com/hardikn13/book-management-system/pkg/controllers"
)

func CollectionOfRoutes(r *mux.Router) {
	r.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
