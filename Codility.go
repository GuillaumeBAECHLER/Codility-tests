package main

import (
	"fmt"

	timecomplexity "./3_TimeComplexity"
)

func main() {
	fmt.Println(timecomplexity.TapeEquilibrium([]int{1, 2, 4, 5}))
	fmt.Println(timecomplexity.TapeEquilibrium([]int{3, 1, 2, 4, 3}))
	fmt.Println(timecomplexity.TapeEquilibrium([]int{100, 13, 20, -10, 4, -5, 2, 3, 4, 5, -10, 20, -30, -34, 1, 2, 4, 5, 9}))
	fmt.Println(timecomplexity.TapeEquilibrium([]int{-1000, 1000}))
}
