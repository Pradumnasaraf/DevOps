package main

import (
	"fmt"
	"strconv"
)

func main() {

	age := "25"
	age2 := "25.5"

	convertedAge, _ := strconv.ParseInt(age, 10, 64) // 10 is the base, 64 is the bit size . We ignore the error by using _.
	fmt.Println(convertedAge +2)

	convertedAge2, _ := strconv.ParseFloat(age2, 64)
	fmt.Println(convertedAge2 + 2.6)
}
