package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://pradumnasaraf.dev:3000/services?service=twitter"

func main() {

	// Parsing the URL
	results, _ := url.Parse(myURL)
	fmt.Println(results.Scheme)   // Scheme is the protocol used. Eg: http, https, ftp, etc
	fmt.Println(results.Host)     // Host is the domain name
	fmt.Println(results.Port())   // Port is the port number
	fmt.Println(results.RawQuery) // RawQuery is the query string
	fmt.Println(results.Query())  // Query is the query string in map form

	data := results.Query()
	fmt.Println(data["service"])

	// &url.URL{} is a pointer to a url.URL type
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "pradumnasaraf.dev",
		Path:     "/services",
		RawQuery: "user=pradumna",
	}

	// String() method is used to convert the url.URL type to string
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
