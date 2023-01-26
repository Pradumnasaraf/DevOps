package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"error"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	quantity int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Book 1", Author: "Author 1", quantity: 10},
	{ID: "2", Title: "Book 2", Author: "Author 2", quantity: 20},
	{ID: "3", Title: "Book 3", Author: "Author 3", quantity: 30}
}


func main(){
	http.HandleFunc("/hello-world",  func( w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe(":8080", nil)
}