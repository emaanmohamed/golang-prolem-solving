package main

//Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
//
//Each letter in magazine can only be used once in ransomNote.
//
//
//
//Example 1:
//
//Input: ransomNote = "a", magazine = "b"
//Output: false
//Example 2:
//
//Input: ransomNote = "aa", magazine = "ab"
//Output: false
//Example 3:
//
//Input: ransomNote = "aa", magazine = "aab"
//Output: true
//

func canConstruct(ransomNote string, magazine string) bool {
	hashmap := make(map[rune]int)
	for _, v := range magazine {
		hashmap[v]++
	}

	for _, val := range ransomNote {
		if hashmap[val] == 0 {
			return false
		} else {
			hashmap[val]--
		}
	}
	return true
}
