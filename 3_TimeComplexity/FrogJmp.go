package timecomplexity

import "math"

// FrogJmp : Gives the number of jumps needed to overgo to a Y distance, given X the starting point and D the distance of each jump
func FrogJmp(X int, Y int, D int) int {
	return int(math.Ceil(float64(Y-X) / float64(D)))
}
