package solution

import "testing"

func TestBinaryGap(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example test n=1041=10000010001", args{1040}, 5},
		{"example test n=15=1111_2", args{15}, 0},
		{"example test n=32=100000_2", args{32}, 0},
		{"extremes", args{1}, 0},
		{"extremes", args{5}, 1},
		{"extremes", args{2147483647}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryGap(tt.args.N); got != tt.want {
				t.Errorf("BinaryGap(%v) = %v, want %v", tt.args.N, got, tt.want)
			}
		})
	}
}
