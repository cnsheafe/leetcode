package two_sum

import (
	"reflect"
	"testing"
)

type testInput struct {
	nums     []int
	target   int
	expected []int
}

var inputs = []testInput{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSumBrute(t *testing.T) {
	for _, input := range inputs {
		if indices := TwoSumBrute(input.nums, input.target); !reflect.DeepEqual(indices, input.expected) {
			t.Errorf("Expected [%d, %d] but got [%d, %d]",
				input.expected[0], input.expected[1],
				indices[0], indices[1])

		}
	}
}

func TestTwoSum(t *testing.T) {
	for _, input := range inputs {
		if indices := TwoSum(input.nums, input.target); !reflect.DeepEqual(indices, input.expected) {
			t.Errorf("Expected [%d, %d] but got [%d, %d]",
				input.expected[0], input.expected[1],
				indices[0], indices[1])

		}
	}
}
