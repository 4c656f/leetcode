package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    tests := []struct {
        name   string
        arg1 []int
		arg2 [][]int
		
        want int64
    }{
        {"base case", []int{1,0,4,6}, [][]int{{2,1},{6,3}}, 9},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := minimumTotalDistance(tt.arg1, tt.arg2)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("minimum_total_distance_traveled(%v, %v) = %v, want %v", tt.arg1, tt.arg2, got, tt.want)
            }
        })
    }
    }
