package arrays

import "testing"

func TestOddOccurrencesInArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1, 2, 3, 2, 1}}, 3},
		{"test 2", args{[]int{20, 18, 100, 45, 30, 29, 18, 100, 100, 45, 20, 10, 30, 100, 29}}, 10},
		{"test 3", args{[]int{20000, 20000, 300000, 20000, 1000000, 1000000}}, 20000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddOccurrencesInArray(tt.args.A); got != tt.want {
				t.Errorf("OddOccurrencesInArray(%v) = %v, want %v", tt.args.A, got, tt.want)
			}
		})
	}
}
