package main

func longestCommonPrefix(strs []string) string {
	firstStr := strs[0]

	for _, val := range strs {
		i := 0
		for i < len(val) && i < len(firstStr) && firstStr[i] == val[i] {
			i++
		}
		firstStr = firstStr[:i]

	}
	return firstStr

}

func main() {
	// Example usage
	input := []string{"flower", "flow", "flight"}
	// Output: fl
	println(longestCommonPrefix(input))
}
