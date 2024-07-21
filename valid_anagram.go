package main

import (
	"sort"
	"strings"
)

//Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
//
//
//Example 1:
//
//Input: s = "anagram", t = "nagaram"
//Output: true
//Example 2:
//
//Input: s = "rat", t = "car"
//Output: false

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	str1 := strings.Split(s, "")
	str2 := strings.Split(t, "")

	sort.Strings(str1)
	sort.Strings(str2)

	return strings.Join(str1, "") == strings.Join(str2, "")
}
