package sorting

import "fmt"

// bubbleSort sorts a slice of integers in ascending order using the bubble sort algorithm.
// Time complexity: O(n^2), Space complexity: O(1).
func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				fmt.Println(nums)
			}
		}
	}
	return nums
}
