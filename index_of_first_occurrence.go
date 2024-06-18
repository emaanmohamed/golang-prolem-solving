package main

import "strings"

func strStr(haystack string, needle string) int {
	if !strings.Contains(haystack, needle) {
		return -1
	} else {
		for index, val := range haystack {
			if uint8(val) == needle[0] {
				return index
				break
			}
		}
	}
	return -1
}

func main() {

}
