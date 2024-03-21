package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Enter a Number:")
	fmt.Scan(&num1)
	if num1%2 == 0 {
		fmt.Print("Entered Number is even")
	} else {
		fmt.Print("Entered NUmber is odd")
	}
}
