package main

import (
	"testing"
)

func BenchmarkRunningSumOf1dArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{1, 2, 3, 4}
		runningSum(nums)
	}
}
func BenchmarkRunningSumOf1dArray1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{1, 2, 3, 4}
		runningSum1(nums)
	}
}
func BenchmarkRunningSumOf1dArray2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{1, 2, 3, 4}
		runningSum2(nums)
	}
}
