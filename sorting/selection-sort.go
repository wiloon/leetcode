package sorting

import "fmt"

// selectionSort sorts a slice of integers in ascending order using the selection sort algorithm.
// Time complexity: O(n^2), Space complexity: O(1).
func selectionSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
		fmt.Println(nums)
	}
	return nums
}
