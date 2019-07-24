package timecomplexity

import (
	"math"
)

// TapeEquilibrium : eturns the minimal difference that can be achieved in an array
func TapeEquilibrium(A []int) int {
	sum := 0
	acc := 0
	minimalDiff := math.MaxFloat64
	for _, number := range A {
		sum += number
	}
	for i := 0; i < (len(A) - 1); i++ {
		acc += A[i]
		minimalDiff = math.Min(minimalDiff, math.Abs((float64(acc) - float64(sum-acc))))
	}
	return int(minimalDiff)
}
