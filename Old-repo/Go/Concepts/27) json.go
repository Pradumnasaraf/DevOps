package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"coursename"`
	Age  int    `json:"age"`
	DOBY int    `json:"doby"`
}

func main() {

	// structToJson()
	// jsonToStruct()
	extractJsonData()

}

func structToJson() {
	tom := []Person{
		{
			Name: "Tom",
			Age:  21,
			DOBY: 199,
		},
		{
			Name: "Ben",
			Age:  21,
			DOBY: 199,
		},
		{
			Name: "Loft",
			Age:  21,
			DOBY: 199,
		},
	}

	// Marshal will convert the struct to json
	jsonOutput, err := json.MarshalIndent(tom, "", "  ")
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(string(jsonOutput))
}

func jsonToStruct() {
	jsonData := []byte(`{
    "coursename": "ben",
    "age": 200,
    "doby": 1975
}`)

	var myUser Person

	validateJson := json.Valid(jsonData)
	if validateJson == true {
		// Unmarshal will convert the json to struct
		json.Unmarshal(jsonData, &myUser)
		fmt.Println(myUser)
	} else {
		fmt.Println("JSON is invalid")
	}

}

func extractJsonData() {

	jsonData := []byte(`
		{
			"coursename": "Mern",
			"price": 200,
			"website": "yt",
			"tags": [ "full-stack", "js" ]
		}
	`)

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData) //

	// We can use range to loop through the map and print the key and value
	for k, v := range myData {
		fmt.Printf("Key is %v with the value %v and type is: %T\n", k, v, v)
	}

}
