package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)
	fmt.Println("Fibonacci series up to", num, ":", fibonacci(num))
}
