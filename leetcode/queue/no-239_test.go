package queue

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"-2",
			args{
				nums: []int{7, 2, 4},
				k:    2,
			},
			[]int{7, 4},
		},
		{
			"-1",
			args{
				nums: []int{1, -1},
				k:    1,
			},
			[]int{1, -1},
		},
		{
			"0",
			args{
				nums: []int{1},
				k:    1,
			},
			[]int{1},
		},
		{
			"1",
			args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			[]int{3, 3, 5, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
			if got := maxSlidingWindowIteration(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
			if got := maxSlidingWindowDP(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
