package main

import "fmt"

func countNum(num uint32) int32 {
	count := int32(0)
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			count++
		}
		// right shift
		num >>= 1
	}
	return count
}

func main() {
	number := uint32(0xFFFFFFFF)

	setBitsCount := countNum(number)
	fmt.Printf("Number of set bits in %d", setBitsCount) // Number of set bits in 32

}
