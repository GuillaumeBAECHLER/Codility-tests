package countingelements

// PermCheck : tells wether an array is a permutation or not
func PermCheck(A []int) int {
	n := len(A)
	tmp := make([]int, n+1)
	for _, number := range A {
		if number <= n && tmp[number] != 1 {
			tmp[number]++
		} else {
			return 0
		}
	}
	return 1
}
