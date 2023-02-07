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

	//encodeJson()
	//decodeJson()
	extractJsonData()


}

func encodeJson() {

	cources := []course{
		{"React", 100, "yt", "nn1234", []string{"react", "js"}},
		{"Mern", 200, "yt", "fgf1234", []string{"full-stack", "js"}},
		{"Golang", 300, "yt", "jh1234", nil},
	}

	finalJson, err := json.MarshalIndent(cources, "", "\t") // json.Marshal() will return byte array

	if err != nil {
		panic(err)
	}

	fmt.Print(string(finalJson))
	// or
	fmt.Printf("%s",finalJson)
}

func decodeJson() {

	jsonData := []byte(`
		{
			"coursename": "Mern",
			"price": 200,
			"website": "yt",
			"tags": [ "full-stack", "js" ]
		},
	`)

	var myCourse course 
	
	checkValid := json.Valid(jsonData)
	if checkValid {
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else{
		fmt.Println("JSON NOT VALID")
	}

}


func extractJsonData(){

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

	for k, v := range myData{
		fmt.Printf("Key is %v with the value %v and type is: %T\n", k, v, v)
	}

}