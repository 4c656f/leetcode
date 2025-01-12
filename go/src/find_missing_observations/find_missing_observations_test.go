package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		arg1 []int
		arg2 int
		arg3 int

		want []int
	}{
		{"base case", []int{3, 2, 4, 3}, 4, 2, []int{6, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingRolls(tt.arg1, tt.arg2, tt.arg3)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find_missing_observations(%v, %v, %v) = %v, want %v", tt.arg1, tt.arg2, tt.arg3, got, tt.want)
			}
		})
	}
}
