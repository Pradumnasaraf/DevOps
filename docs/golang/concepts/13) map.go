package main

import (
	"fmt"
)

func main() {
	var languages = map[string]string{}
	// var languages = make(map[string]string, 0) // create a slice of maps with a length of 0. // map[keyType]valueType  

	//laguages[key] = value
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Go"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println(languages["JS"]) // access the value of a key

	delete(languages, "RB") // Deleting a Key (No key means no Value as well)
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println(key, value)
	}
}
