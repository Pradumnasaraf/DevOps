package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("First Statement")
	go greeter("Second Statement")
	greeter("Third Statement")

}

func greeter(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
