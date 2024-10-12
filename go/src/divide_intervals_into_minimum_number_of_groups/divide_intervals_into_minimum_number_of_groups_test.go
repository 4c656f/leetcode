package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		arg1 [][]int

		want int
	}{
		{"base case", [][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minGroups(tt.arg1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divide_intervals_into_minimum_number_of_groups(%v) = %v, want %v", tt.arg1, got, tt.want)
			}
		})
	}
}
