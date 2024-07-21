package main

import "fmt"

//write a function that takes a string and return the number of each character occurance

func countDuplication(str string) map[rune]int {
	hashmap := make(map[rune]int)

	for _, char := range str {
		hashmap[char]++
	}
	return hashmap
}

func main() {
	// Example usage
	input := "hello"
	result := countDuplication(input)
	// Output: map[h:1 e:1 l:2 o:1]
	for key, value := range result {
		println(string(key), ":", value)
	}
}

func slicer() {
	data := []int{1, 2, 3, 4, 5, 6}
	res := data[:3]
	final := res[2:5]
	fmt.Println(final)
}
