package main

import (
	"fmt"
)

// jwtToken := 30000 // is not allowed outside the main function
//  var jwtToken int = 3000 // is allowed

const LoginToken string = "sbdghsghjsd" // geberish values

// to check its string type
func main() {
	var username string = "Samiul"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	//to check if it is boolean type
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	//uint is
	var smallVal uint8 = 254
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 254.3843743
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//default values and some alices
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T\n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style: we can totally ignore var while declaring variable
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T\n", LoginToken)

}
