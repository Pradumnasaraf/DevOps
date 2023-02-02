package main

import (
	"fmt"
	"strconv"
)

func main() {
	var bookings = make([]map[string]string, 0) // make() is used to create a slice, map, or chan with initial values

	for i := 0; i < 1; i++ {

		name := "John"
		age := 32
		city := "New York"

		var myMap = make(map[string]string)

		myMap["name"] = name
		myMap["age"] = strconv.FormatInt(int64(age), 10)
		myMap["city"] = city

		bookings = append(bookings, myMap)
		bookings = append(bookings, myMap)
		bookings = append(bookings, myMap)

		fmt.Println(bookings)
		for _ , booking := range bookings {
			fmt.Println(booking["name"]) // access the value of a key
		}
	}
}
