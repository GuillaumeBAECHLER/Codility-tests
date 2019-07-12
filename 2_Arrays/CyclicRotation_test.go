package arrays

import (
	"reflect"
	"testing"
)

func TestCyclicRotation(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{[]int{1, 2, 3, 4}, 1}, []int{4, 1, 2, 3}},
		{"test 2", args{[]int{1, 2, 3, 4, 5, 6}, 9}, []int{4, 5, 6, 1, 2, 3}},
		{"test 3", args{[]int{}, 3}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CyclicRotation(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CyclicRotation(%v, %v) = %v, want %v", tt.args.A, tt.args.K, got, tt.want)
			}
		})
	}
}
