package main

//Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3,4,5,6,7], k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]

func rotate(nums []int, k int) {
	temp := make([]int, 0, len(nums))
	temp = append(temp, nums[len(nums)-k:]...)
	temp = append(temp, nums[:len(nums)-k]...)
	copy(nums, temp)
}


