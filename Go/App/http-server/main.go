package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 10},
	{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 20},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: 30},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}
func createBooks(c *gin.Context){
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func root (c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "API is working fine",
    })
}

func main(){
	router := gin.Default()
	router.GET("/", root)
	router.GET("/books", getBooks)
	router.POST("/books", createBooks)
	router.Run("localhost:8083")
}