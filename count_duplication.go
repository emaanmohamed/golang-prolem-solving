package main

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
