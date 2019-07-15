package timecomplexity

import "math"

// FrogJmp
func FrogJmp(X int, Y int, D int) int {
	return int(math.Ceil(float64(Y-X) / float64(D)))
}
