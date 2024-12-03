package Binary_Search

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := (left + right) / 2
		if target == matrix[mid/n][mid%n] {
			return true
		} else if target > matrix[mid/n][mid%n] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false

}
