package main

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func main() {
	// Example usage
	input1 := "acb"
	input2 := "ahbgdc"
	// Output: true
	println(isSubsequence(input1, input2))
	input3 := "axc"
	input4 := "ahbgdc"
	// Output: false
	println(isSubsequence(input3, input4))
}
