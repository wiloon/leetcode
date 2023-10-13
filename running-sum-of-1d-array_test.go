package main

import (
	"fmt"
	"testing"
)

func validate(actual, expected []int) {
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			fmt.Println("Fail")
			return
		}
	}
	fmt.Println("Pass")
}

func TestCase1(t *testing.T) {
	in := []int{1, 2, 3, 4}
	out := []int{1, 3, 6, 10}

	var tmp []int
	tmp = make([]int, len(in))
	copy(tmp, in)
	result := runningSum(tmp)
	validate(result, out)

	tmp = make([]int, len(in))
	copy(tmp, in)
	result1 := runningSum1(tmp)
	validate(result1, out)

	tmp = make([]int, len(in))
	copy(tmp, in)
	result2 := runningSum2(tmp)
	validate(result2, out)

	tmp = make([]int, len(in))
	copy(tmp, in)
	result3 := runningSum3(tmp)
	validate(result3, out)
}

func TestCase2(t *testing.T) {
	in := []int{1, 1, 1, 1, 1}
	out := []int{1, 2, 3, 4, 5}
	result := runningSum(in)
	validate(result, out)
}

func TestCase3(t *testing.T) {
	nums := []int{3, 1, 2, 10, 1}
	out := []int{3, 4, 6, 16, 17}
	result := runningSum(nums)
	validate(result, out)
}
