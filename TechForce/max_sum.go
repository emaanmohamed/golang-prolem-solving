package main

import (
	"fmt"
)

func findMaxSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	firstMax, secondMax := arr[0], arr[1]

	if secondMax > firstMax {
		firstMax, secondMax = secondMax, firstMax
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] > firstMax {
			secondMax = firstMax
			firstMax = arr[i]
		} else if arr[i] > secondMax {
			secondMax = arr[i]
		}
	}

	return firstMax + secondMax
}

func main() {
	list := []int{5, 9, 7, 11}
	fmt.Println(findMaxSum(list))
}
