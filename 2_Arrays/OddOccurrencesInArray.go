package arrays

// OddOccurrencesInArray : returns the unpaired element of an array
func OddOccurrencesInArray(A []int) int {
	unpaired := make(map[int]bool)
	for _, number := range A {
		if !unpaired[number] {
			unpaired[number] = true
		} else {
			delete(unpaired, number)
		}
	}
	// unpaired should have a size of 1 at the end
	var unpairedNumber int
	for k := range unpaired {
		unpairedNumber = k
	}
	return unpairedNumber
}
