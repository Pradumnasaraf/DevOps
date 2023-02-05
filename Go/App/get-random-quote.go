package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	Id     int    `json:"_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Tags   []string `json:"tags"`
}

func main() {

	url := "https://api.quotable.io/random"

	res, err := http.Get(url)
	checkNilError(err)

	content, err := ioutil.ReadAll(res.Body)
	checkNilError(err)

	defer res.Body.Close()

	quoteData := Quote{} 
	json.Unmarshal(content, &quoteData)
	
	fmt.Println(quoteData.Content)
	fmt.Println(quoteData.Author)
	
	for _, tag := range quoteData.Tags {
		fmt.Println(tag)
	}

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
