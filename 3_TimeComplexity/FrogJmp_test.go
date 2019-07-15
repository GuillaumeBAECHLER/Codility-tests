package timecomplexity

import "testing"

func TestFrogJmp(t *testing.T) {
	type args struct {
		X int
		Y int
		D int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{10, 85, 30}, 3},
		{"test 2", args{0, 100, 20}, 5},
		{"test 3", args{20, 100, 20}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrogJmp(tt.args.X, tt.args.Y, tt.args.D); got != tt.want {
				t.Errorf("FrogJmp(%v, %v, %v) = %v, want %v", tt.args.X, tt.args.Y, tt.args.D, got, tt.want)
			}
		})
	}
}
