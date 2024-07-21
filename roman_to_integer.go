package main

func romanToInt(s string) int {
	sum := 0
	hash := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, v := range s {
		sum += hash[string(v)]
		if i != 0 {
			if hash[string(v)] > hash[string(s[i-1])] {
				sum -= 2 * hash[string(s[i-1])]
			}
		}

	}

	return sum
}

func main() {
	// Example usage
	input := "III"
	// Output: 3
	println(romanToInt(input))
	input2 := "IV"
	// Output: 4
	println(romanToInt(input2))
	input3 := "IX"
	// Output: 9
	println(romanToInt(input3))
	input4 := "LVIII"
	// Output: 58
	println(romanToInt(input4))
	input5 := "MCMXCIV"
	// Output: 1994
	println(romanToInt(input5))
}
