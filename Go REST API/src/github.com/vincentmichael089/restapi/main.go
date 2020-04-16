//Tutorial by TraversyMedia https://www.youtube.com/watch?v=SonwZ6MF5BE
package main

import (
	//"encoding/json"
	//"log"
	"net/http"
	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

//book model
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

//author model
type Author struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

//init books var as a slice of book struct
var books []Book

//functions
func getBooks(w http.ResponseWriter, r *http.Request){

}

func getBook(w http.ResponseWriter, r *http.Request){
	
}

func createBook(w http.ResponseWriter, r *http.Request){
	
}

func updateBook(w http.ResponseWriter, r *http.Request){
	
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	
}


func main() {
	//initialize router
	router := mux.NewRouter()

	//mock data
	books = append(books, Book{
		ID : "1",
		Isbn : "123123123",
		Title : "Booky McBookface",
		Author : &Author{
			FirstName : "Boo",
			LastName : "Baa"}})

	books = append(books, Book{
		ID : "2",
		Isbn : "2331223",
		Title : "Golang 101",
		Author : &Author{
			FirstName : "Go",
			LastName : "Lang"}})

	//route handlers / endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//run server
	http.ListenAndServe(":8000", router) 
}