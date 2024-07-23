package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza app")
	fmt.Println("Plesae rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin) // create a reader object
	intput, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating our pizza", intput)

	numRating, err := strconv.ParseInt(strings.TrimSpace(intput), 10, 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Thanks for rating our pizza, add 1 to your ratting", numRating+1)
	}
}
