package main

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}

func runningSum1(nums []int) []int {
	l := len(nums) + 1
	out := make([]int, l)
	for i := 1; i < l; i++ {
		out[i] = out[i-1] + nums[i-1]
	}
	return out[1:]
}

func runningSum2(nums []int) []int {
	out := make([]int, len(nums)+1)
	for i, v := range nums {
		out[i+1] = out[i] + v
	}
	return out[1:]
}
func runningSum3(nums []int) []int {
	tmp := 0
	out := make([]int, len(nums))
	for i, v := range nums {
		tmp = tmp + v
		out[i] = tmp
	}
	return out
}
