package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 { // Here 3 will not be printed
			continue
		}
		if i == 4 { // From here the loop will get exited without printing 4
			break
		}
		fmt.Println(i) // only 0, 1, 2 will be printed
	}
}
