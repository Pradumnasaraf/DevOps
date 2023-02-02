package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://pradumnasaraf.dev"

func main() {

	res, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Response is of type: %T\n", res)
	fmt.Println(res.Status)

	databytes, err := ioutil.ReadAll(res.Body) // Reading the response in desired format. We can't read the response directly.
	checkNilError(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
