package arrays

// CyclicRotation : rotates an array A to the right K times
func CyclicRotation(A []int, K int) []int {
	B := make([]int, len(A))
	if len(A) == 0 {
		return B
	}
	shift := (K % len(A))
	for i := range A {
		B[(i+shift)%len(A)] = A[i]
	}
	return B
}
