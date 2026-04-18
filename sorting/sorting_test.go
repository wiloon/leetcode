package sorting

import (
	"reflect"
	"testing"
)

var sortTestCases = []struct {
	name  string
	input []int
	want  []int
}{
	{
		name:  "unsorted slice",
		input: []int{5, 3, 8, 1, 2},
		want:  []int{1, 2, 3, 5, 8},
	},
	{
		name:  "already sorted",
		input: []int{1, 2, 3, 4, 5},
		want:  []int{1, 2, 3, 4, 5},
	},
	{
		name:  "reverse sorted",
		input: []int{5, 4, 3, 2, 1},
		want:  []int{1, 2, 3, 4, 5},
	},
	{
		name:  "single element",
		input: []int{42},
		want:  []int{42},
	},
	{
		name:  "duplicate elements",
		input: []int{3, 1, 2, 1, 3},
		want:  []int{1, 1, 2, 3, 3},
	},
	{
		name:  "empty slice",
		input: []int{},
		want:  []int{},
	},
}

func TestSortingAlgorithms(t *testing.T) {
	algorithms := []struct {
		name string
		fn   func([]int) []int
	}{
		{"bubbleSort", bubbleSort},
		{"selectionSort", selectionSort},
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tt := range sortTestCases {
				t.Run(tt.name, func(t *testing.T) {
					input := make([]int, len(tt.input))
					copy(input, tt.input)
					result := algo.fn(input)
					if !reflect.DeepEqual(result, tt.want) {
						t.Errorf("%s() = %v, want %v", algo.name, result, tt.want)
					}
				})
			}
		})
	}
}
