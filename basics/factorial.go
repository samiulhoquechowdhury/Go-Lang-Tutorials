package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int
	fmt.Println("Enter a Number: ")
	fmt.Scanln(&num)
	fmt.Println("Factorial:", factorial(num))
}
