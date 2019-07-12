package main

import (
	arrays "./2_Arrays"
)

func main() {
	arrays.CyclicRotation([]int{1, 2, 3, 4, 5, 6}, 1)
	arrays.CyclicRotation([]int{1, 2, 3, 4, 5, 6}, 2)
	arrays.CyclicRotation([]int{1, 2, 3, 4, 5, 6}, 3)
	arrays.CyclicRotation([]int{1, 2, 3, 4, 5, 6}, 4)
}
