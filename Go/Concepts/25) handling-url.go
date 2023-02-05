package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://pradumnasaraf.dev:3000/services?service=twitter"

func main() {

	fmt.Println(myURL)

	// Parsing

	results, _ := url.Parse(myURL)
	fmt.Println(results.Scheme)
	fmt.Println(results.Host)
	fmt.Println(results.Port())
	fmt.Println(results.RawQuery)
	fmt.Println(results.Query())

	data := results.Query()
	fmt.Println(data["service"])

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "pradumnasaraf.dev",
		Path: "/services",
		RawQuery: "user=pradumna",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
