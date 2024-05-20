package main

import "math"

//If num = 36, the square root is 6.
//The factors of 36 are 2, 3, 4, 6, 9, 12, 18.
//Notice that factors come in pairs that multiply to num: (2, 18), (3, 12), (4, 9), (6, 6).

func countPrimes(n int) int {
	if n > int(1)<<40 {
		panic("n exceeds maximum number (1<<40).")
	}
	count := 0
	for i := 0; i < n; i++ {
		if IsPrime(i) {
			count++
		}
	}
	return count
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for j := 2; j <= int(math.Sqrt(float64(n))); j++ {
		if n%j == 0 {
			return false
		}
	}
	return true

}

func main() {
	n := 36
	count := countPrimes(n)
	println(count)
}
