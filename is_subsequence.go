package main

import "strings"

func isSubsequence(s string, t string) bool {
	arr := make(map[int]string)
	hashmap := make(map[int]string)
	for _, val := range s {
		arr[] = string(val)
		if strings.Contains(t, string(val)) {
			hashmap[] = string(val)
		} else {
			return false
		}
	}
	for i := range arr {
		if arr[i] != hashmap[i] {
			return false
		}
	}
	return true
}
