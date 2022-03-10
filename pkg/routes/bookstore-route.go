package routes

import (
	"github.com/gorilla/mux"
	"github.com/ashraf-mse/book-store/pkg/controllers"
)


var RegisterBookStoreRoutes = func(router *mux.Router){
	//Get all books
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")

	//Get book by ID
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")

	//Create a new book 
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")

	//Update a book
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")

	//Delete a book
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}