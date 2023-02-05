package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// getRequest()
	// postRequest()
	postFormRequest()

}

func getRequest() {
	const apiUrl = "https://ossapi.vercel.app/api/devtip"

	res, err := http.Get(apiUrl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Printf("Status code: %v\n", res.StatusCode)

	var resString strings.Builder

	// normal string method
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))

	// Using the strings library
	byteCount, _ := resString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(resString.String())

}

func postRequest() {

	const apiUrl = "https://ossapi.vercel.app/api/devtip"

	// fake json paylaod

	reqBody := strings.NewReader(`
		{ 
			"message": "eat good food"	
		}
	`)

	res, err := http.Post(apiUrl, "application/json" , reqBody)

	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func postFormRequest(){

	const apiUrl = "https://ossapi.vercel.app/api/devtip"

	data := url.Values{}
	data.Add("message", "Hey, drink water daily")

	res, err := http.PostForm(apiUrl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
