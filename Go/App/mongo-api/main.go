package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/pradumnasaraf/mongo-api/router"
)

func main(){
	r := router.Router()
	fmt.Println("Server is started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Print("Listening at port: 4000")
}