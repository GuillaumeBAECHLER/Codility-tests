package countingelements

import "testing"

func TestPermCheck(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1, 2, 3, 4, 5}}, 1},
		{"test 2", args{[]int{1, 2, 3, 4, 6}}, 0},
		{"test 3", args{[]int{4, 2, 3, 1, 5}}, 1},
		{"test 4", args{[]int{1, 2, 2, 4, 5, 6, 7}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermCheck(tt.args.A); got != tt.want {
				t.Errorf("PermCheck(%v) = %v, want %v", tt.args.A, got, tt.want)
			}
		})
	}
}
