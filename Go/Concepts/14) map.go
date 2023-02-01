package main

import (
	"fmt"
)

func main() {
	var languages = make(map[string]string, 0) // create a slice of maps

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Go"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println(languages["JS"]) // access the value of a key

	delete(languages, "RB") // delete a key from a map
	fmt.Println(languages)

	for _, value := range languages {
		fmt.Println(value)
	}
}
