package main

import (
	"fmt"
	"strings"
)

//Problem 01
//A palindrome is a word or sequence of characters which reads the same backward or
//forward. For example the words: level, anna, noon, rotator are all palindromes.
//Write a function palindrom that accepts a string as an argument and returns a boolean
//indicating whether the input is a palindrome or not, for example:
//        palindrome("anna") # returns True
//palindrome("apple") # returns False

func palindrome(str string) bool {
	str = strings.ToLower(str)
	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func main() {
	fmt.Println(palindrome("anna"))
	fmt.Println(palindrome("apple"))

}
