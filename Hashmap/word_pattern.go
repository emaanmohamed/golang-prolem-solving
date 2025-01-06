package main

import "strings"

//Given a pattern and a string s, find if s follows the same pattern.
//
//Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:
//
//Each letter in pattern maps to exactly one unique word in s.
//Each unique word in s maps to exactly one letter in pattern.
//No two letters map to the same word, and no two words map to the same letter.
//
//
//Example 1:
//
//Input: pattern = "abba", s = "dog cat cat dog"
//
//Output: true
//
//Explanation:
//
//The bijection can be established as:
//
//'a' maps to "dog".
//'b' maps to "cat".
//Example 2:
//
//Input: pattern = "abba", s = "dog cat cat fish"
//
//Output: false
//
//Example 3:
//
//Input: pattern = "aaaa", s = "dog cat cat dog"
//
//Output: false

func wordPattern(pattern string, s string) bool {

	hashmap1 := make(map[string]int)
	hashmap2 := make(map[string]int)

	arr := strings.Split(s, " ")

	if len(pattern) != len(arr) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		char := string(pattern[i])
		if hashmap1[char] != hashmap2[arr[i]] {
			return false
		}
		hashmap1[char] = i + 1
		hashmap2[arr[i]] = i + 1
	}
	return true
}
