package main

// Write a function to reverse a string in Go

func reverse_string(str string) string {

	runes := []rune(str)
	start := 0
	end := len(runes) - 1
	// Swap the runes in place.
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	// Convert the slice of runes back to a string.
	return string(runes)
}

func main() {

}
