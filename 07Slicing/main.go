package main

import "fmt"

func main() {
	fmt.Println("welcome to slices in golang")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fuitlist is %T", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList) 

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

}
