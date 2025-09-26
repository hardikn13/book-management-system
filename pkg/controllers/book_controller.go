package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hardikn13/book-management-system/pkg/models"
	"github.com/hardikn13/book-management-system/pkg/utils"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.FindAllBooks()
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	book := newBook.AddBook()
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("unable to parse the bookId")
	}
	w.Header().Set("Content-Type", "application/json")
	book, _ := models.FindBook(ID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("unable to parse the bookId")
	}
	bookDetails, db := models.FindBook(ID)
	if newBook.Name != "" {
		bookDetails.Name = newBook.Name
	}
	if newBook.Author != "" {
		bookDetails.Author = newBook.Author
	}
	if newBook.Publication != "" {
		bookDetails.Publication = newBook.Publication
	}
	db.Save(&bookDetails)
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("unable to parse the bookId")
	}
	w.Header().Set("Content-Type", "application/json")
	book := models.RemoveBook(ID)
	res, _ := json.Marshal(book)
	w.Write(res)
}
