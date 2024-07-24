package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()
	fmt.Println("Present Time is: ", presentTime)
	fmt.Println("Present Time is: ", presentTime.Format("01-02-2006 Mon"))
	
	createdDate := time.Date(2020, time.August, 10, 23,23,0,0 , time.UTC)
	fmt.Println("Created Date is: ", createdDate)
	fmt.Println("Created Date is: ", createdDate.Format("01-02-2006 Mon")) 
}
