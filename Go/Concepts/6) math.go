package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano()) // seed the random number generator. Time is used as seed to generate random numbers
// 	fmt.Println(rand.Intn(100) +1) // generate random number between 1 and 100
// }

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(100)) // generate random number between 1 and 100
	fmt.Println(myRandomNumber)
}
