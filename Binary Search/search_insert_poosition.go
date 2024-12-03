package Binary_Search

//Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
//You must write an algorithm with O(log n) runtime complexity.
//
//
//
//Example 1:
//
//Input: nums = [1,3,5,6], target = 5
//Output: 2
//Example 2:
//
//Input: nums = [1,3,5,6], target = 2
//Output: 1
//Example 3:
//
//Input: nums = [1,3,5,6], target = 7
//Output: 4

func searchInsert(nums []int, target int) int {
	first := 0
	last := len(nums) - 1

	for first < last {
		mid := (first + last) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			first = mid
		} else {
			last = mid
		}
	}
	return first + 1

}

// or

//func searchInsert(nums []int, target int) int {
//	left, right := 0, len(nums) - 1
//	for right >= left {
//		mid := (left+right)/2
//		if nums[mid] == target {
//			return mid
//		}
//		if target > nums[mid] {
//			left = mid + 1
//		} else {
//			right = mid - 1
//		}
//	}
//	return left
//}
