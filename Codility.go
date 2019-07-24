package main

import (
	"fmt"

	timecomplexity "./3_TimeComplexity"
)

func main() {
	fmt.Println(timecomplexity.PermMissingElem([]int{1, 2, 4, 5}))
	fmt.Println(timecomplexity.PermMissingElem([]int{}))
}
