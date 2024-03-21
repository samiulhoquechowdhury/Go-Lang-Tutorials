package main

import (
	"fmt"
)

// to check its string type
func main() {
	var username string = "Samiul"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	//to chekc if it is boolean type
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	//
	var smallVal uint8 = 254
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)
}
