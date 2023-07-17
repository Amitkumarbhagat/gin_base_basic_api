package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Amit Bhagat", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "Sumit bhagat", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Ankit bhagat", Quantity: 6},
}

// the gin context is all the information about the request, and
//it allows you to return a response

// by indentedJson we can return file text etc. in this case returning
// book slice in well manner
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	//	router responsible for handle different kind of routes and kind
	//	of different end points of our api
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
