package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Print("Server is running on port 8080 \n")
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) { // w is used to write the response and r is used to read the request
	w.Write([]byte("Hello World"))
}