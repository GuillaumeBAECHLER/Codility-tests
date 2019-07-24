package timecomplexity

// PermMissingElem : Gives the missing element of an array of N elements storing different integers in the range [1...(N+1)]
func PermMissingElem(A []int) int {
	n := len(A)
	expectedSum := (n + 1) * (n + 2) / 2
	sum := 0
	for _, number := range A {
		sum += number
	}
	return expectedSum - sum
}
