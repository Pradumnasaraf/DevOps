package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const URL = "https://osapi.vercel.app/api/devtip"

func main() {

	// getReq()
	postReq()

}

func getReq() {
	res, err := http.Get(URL)
	checkNilErr(err)

	byteData, err := io.ReadAll(res.Body)
	fmt.Println(string(byteData))
}

func postReq() {
	resBody := strings.NewReader(
		`{
	"message": "Eat good food"
	}`)

	req, err := http.Post(URL, "application/json", resBody)

	checkNilErr(err)
	defer req.Body.Close()

	dataBytes, err := io.ReadAll(req.Body)
	fmt.Println("Status Code", req.StatusCode)
	fmt.Println("Response Body", string(dataBytes))

}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func postFormRequest() {

	const apiUrl = "https://ossapi.vercel.app/api/devtip"

	data := url.Values{}
	data.Add("message", "Hey, drink water daily")

	res, err := http.PostForm(apiUrl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))

}
