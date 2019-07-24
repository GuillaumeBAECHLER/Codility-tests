package timecomplexity

import "testing"

func TestPermMissingElem(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1, 2, 4}}, 3},
		{"test 2", args{[]int{1, 2, 3, 4}}, 5},
		{"test 3", args{[]int{}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermMissingElem(tt.args.A); got != tt.want {
				t.Errorf("PermMissingElem(%v) = %v, want %v", tt.args.A, got, tt.want)
			}
		})
	}
}
