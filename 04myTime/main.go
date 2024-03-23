package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	//this time format is given to extract the original time and day as per the golang documentation
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//tp create a custom date

	createDate := time.Date(2024, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006"))

}
