package main

import "strings"

func strStr(haystack string, needle string) int {
	if !strings.Contains(haystack, needle) {
		return -1
	}

	haystackLen := len(haystack)
	needleLen := len(needle)
	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}

func main() {

}
