package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-bookstore/pkg/models"
	"github.com/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error")
	}
	bookDetails, _ := models.GetBookById(Id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookDetails)

}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error")
	}
	book := models.DeleteBookById(Id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	fmt.Println(string(res))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	params := mux.Vars(r)
	ID, _ := strconv.ParseInt(params["id"], 0, 0)
	bookdetails, db := models.GetBookById(ID)
	if book.Author != "" {
		bookdetails.Author = book.Author
	}
	if book.Name != "" {
		bookdetails.Name = book.Name
	}
	if book.Publication != "" {
		bookdetails.Publication = book.Publication
	}
	db.Save(&bookdetails)
	json.NewEncoder(w).Encode(bookdetails)

}
