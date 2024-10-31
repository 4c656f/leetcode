package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    tests := []struct {
        name   string
        arg1 int
		
        want int
    }{
        {"base case", 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := minimum_number_of_removals_to_make_mountain_array(tt.arg1)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("minimum_number_of_removals_to_make_mountain_array(%v) = %v, want %v", tt.arg1, got, tt.want)
            }
        })
    }
    }

