package main

import "fmt"

func main() {
	// Creating a map where keys are strings and values are integers
	myMap := make(map[string]int)

	// Adding key-value pairs to the map
	myMap["apple"] = 10
	myMap["banana"] = 20
	myMap["orange"] = 15

	// Accessing values using keys
	fmt.Println("Value for key 'apple':", myMap["apple"])
	fmt.Println("Value for key 'banana':", myMap["banana"])

	// Modifying value for an existing key
	myMap["apple"] = 12

	// Deleting a key-value pair
	delete(myMap, "orange")

	// Iterating over the map
	fmt.Println("Printing all key-value pairs:")
	for key, value := range myMap {
		fmt.Println(key, ":", value)
	}

	// Checking if a key exists
	if _, ok := myMap["banana"]; ok {
		fmt.Println("Banana exists in the map")
	} else {
		fmt.Println("Banana does not exist in the map")
	}
}
