package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
	
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serverHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Hey welcome to Golang</h1>"))

}