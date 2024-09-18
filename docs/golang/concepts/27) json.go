package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// It's important to mention `json:"key-name"` otherwise it will pick the keyword mentioned in the Struct with capitalization.
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Place string `json:"place"`
}

func main() {
	// letsMarshal()
	// letsUnMarshal()	
	extractJsonData()

}

func letsUnMarshal() {
	jsonData := []byte(`
		{
			"name": "Pradumna Saraf",
			"age": 25,
			"place": "India"
		}
	`)

	var myData Person

	isJSONValid := json.Valid(jsonData)
	if isJSONValid == true {
		json.Unmarshal(jsonData, &myData)
	} else {
		fmt.Println("Invalid JSON")
	}

	fmt.Println(myData.Age)
	fmt.Println(myData.Name)
	fmt.Println(myData.Place)


}

func letsMarshal() {
	myPerson := Person{
		Name:  "Pradumna Saraf",
		Age:   25,
		Place: "India",
	}

	myPersonSlice := []Person{
		{
			Name:  "Pradumna Saraf",
			Age:   25,
			Place: "India",
		},
		{
			Name:  "Pradumna Saraf",
			Age:   25,
			Place: "India",
		},
		{
			Name:  "Pradumna Saraf",
			Age:   25,
			Place: "India",
		},
	}

	// Single Object
	data, err := json.Marshal(&myPerson)
	if err != nil {
		log.Fatal("Unable to Marshal")
	}
	fmt.Println(string(data))

	// A list of Objects
	sliceData, err := json.MarshalIndent(&myPersonSlice, "", "  ")
	if err != nil {
		log.Fatal("Unable to Marshal")
	}
	fmt.Println(string(sliceData))
}


func extractJsonData() {
	jsonData := []byte(`
		{
			"name": "Pradumna Saraf",
			"age": 25,
			"place": "India"
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