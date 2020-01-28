package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct{
	ID      string  `json := "id"`
	Isbn    string  `json := "isbn"`
	Title   string  `json := "title"`
	Author *Author  `json := "author"`

}


//Author struct
type Author struct{
	Fname      string  `json := "fname"`
	Lname      string  `json := "lname"`
}

//Init books var 
var books []Book

//Get All Books 
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}
//Get single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through books & find with id
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

	
}
//Create a new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	//_ = json.NewDecoder(r.Body).Decode(&book)
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}
//Update a book
func updatBook(w http.ResponseWriter, r *http.Request){
    
}
//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

 func main(){
	 //fmt.Println("Hello world man")
	 //Init router 
	 r := mux.NewRouter()
	 //Mock Data
	 books = append(books,Book{ID: "2", Isbn: "888", Title: "Opekkha", 
			Author:&Author{Fname: "Humayen", Lname: "Ahmed"}})
	books = append(books,Book{ID: "4", Isbn: "8008", Title: "Book one", 
	        Author:&Author{Fname: "Maya", Lname: "Roy"}})
	 //Rout handler / end point
	 r.HandleFunc("/api/books", getBooks).Methods("GET")
	 r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	 r.HandleFunc("/api/book/{id}", createBook).Methods("POST") 
	 r.HandleFunc("/api/book/{id}", updatBook).Methods("PUT")
	 r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	 log.Fatal(http.ListenAndServe(":8000", r))

	  
 }
