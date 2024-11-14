package main

//Given two strings s and t, determine if they are isomorphic.
//
//Two strings s and t are isomorphic if the characters in s can be replaced to get t.

//All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

// egg - app
//
//func isIsomorphic(s string, t string) bool {
//
//	if len(s) != len(t) {
//		return false
//	}
//
//	hashMap := make(map[byte]byte)
//	visited := make(map[byte]bool)
//
//	for i := 0; i < len(s); i++ {
//		if val, ok := hashMap[s[i]]; ok {
//			if val != t[i] {
//				return false
//			}
//		} else {
//			if visited[t[i]] {
//				return false
//			}
//			hashMap[s[i]] = t[i]
//			visited[t[i]] = true
//		}
//	}
//
//	return true
//}

func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}
		sMap[s[i]] = i + 1
		tMap[t[i]] = i + 1
	}
	return true
}

func main() {
	s := "egg"
	t := "app"
	isIsomorphic(s, t)
}
