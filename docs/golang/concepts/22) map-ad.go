package main

import (
	"fmt"
	"strconv"
)

func main() {

	bookings :=[]map[string]string{} // This is a slice of map.

	for i := 0; i < 1; i++ {

		name := "John"
		age := 32
		city := "New York"

		myMap := map[string]string{}

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
