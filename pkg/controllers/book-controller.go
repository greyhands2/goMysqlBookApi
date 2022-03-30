package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorlla/mux"
	"github.com/greyhands2/goMysqlBookApi/pkg/models"
	"github.com/greyhands2/goMysqlBookApi/pkg/utils"
)

var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	result, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)

	res.Write(result)

}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsind")
	}

	bookDetails, _ := models.GetBookById(ID)

	result, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)

	res.Write(result)

}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}

	utils.ParseBody(req, CreateBook)

	b := CreateBook.CreateBook()

	result, _ := json.Marshal(b)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write(result)

}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(ID)

	result, _ := json.Marshal(book)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write(result)

}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var updateBook = &models.Book{}

	utils.ParseBody(req, updateBook)

	vars := mux.Vars(req)

	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	result, _ := json.Marshal(bookDetails)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write(result)
}
