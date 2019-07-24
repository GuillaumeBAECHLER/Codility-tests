package timecomplexity

import "testing"

func TestTapeEquilibrium(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1, 2, 4, 5}}, 2},
		{"test 2", args{[]int{3, 1, 2, 4, 3}}, 1},
		{"test 3", args{[]int{100, 13, 20, -10, 4, -5, 2, 3, 4, 5, -10, 20, -30, -34, 1, 2, 4, 5, 9}}, 61},
		{"test 4", args{[]int{-1000, 1000}}, 2000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TapeEquilibrium(tt.args.A); got != tt.want {
				t.Errorf("TapeEquilibrium(%v) = %v, expected %v", tt.args.A, got, tt.want)
			}
		})
	}
}
