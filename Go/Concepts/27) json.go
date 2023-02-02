package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int	`json:"price"`
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	encodeJson()


}

func encodeJson() {

	cources := []course{
		{"React", 100, "yt", "nn1234", []string{"react", "js"}},
		{"Mern", 200, "yt", "fgf1234", []string{"full-stack", "js"}},
		{"Golang", 300, "yt", "jh1234", nil},
	}

	finalJson, err := json.MarshalIndent(cources, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Print(string(finalJson))
	// or
	fmt.Printf("%s",finalJson)



}
