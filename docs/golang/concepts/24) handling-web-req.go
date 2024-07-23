package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://pradumnasaraf.dev"

func main() {

	res, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Response is of type: %T\n", res)
	fmt.Println(res.Status)

	databytes, err := io.ReadAll(res.Body) //We can't read the response directly. So we use ioutil.ReadAll()
	checkNilError(err)

	content := string(databytes) // convert the byte array to string
	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
