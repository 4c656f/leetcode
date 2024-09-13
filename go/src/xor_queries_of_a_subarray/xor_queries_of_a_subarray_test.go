package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		arg1 []int
		arg2 [][]int

		want []int
	}{
		{"base case", []int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}, []int{2, 7, 14, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := xorQueries(tt.arg1, tt.arg2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xor_queries_of_a_subarray(%v, %v) = %v, want %v", tt.arg1, tt.arg2, got, tt.want)
			}
		})
	}
}
