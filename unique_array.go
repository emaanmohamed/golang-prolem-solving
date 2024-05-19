package main

import "fmt"

//Problem 05
//Write a function unique that takes an array of strings as input and returns an array of
//the unique entries in the input, for example:
//
//unique(['apples', 'oranges', 'flowers', 'apples']) # returns ['oranges', 'flowers']

//unique(['apples', 'apples'])

func unique(arr []string) []string {

	uniqueMap := make(map[string]bool)
	uniqueEntries := []string{}

	for _, item := range arr {

		if !uniqueMap[item] {
			uniqueEntries = append(uniqueEntries, item)
			uniqueMap[item] = true
		}
	}

	return uniqueEntries[1:]
}
func main() {
	// Example usage
	input1 := []string{"apples", "oranges", "flowers", "apples"}
	fmt.Println(unique(input1)) // Output: [apples oranges flowers]

	input2 := []string{"apples", "apples"}
	fmt.Println(unique(input2)) // Output: [apples]
}
