package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Litchi"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list is : ", len(fruitList))

	var vegList = [6]string{"potato", "Tomato", "beans", "Musgrooms"}

	fmt.Println("Vegy list is: ", len(vegList))

}
