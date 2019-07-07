package iterations

import "strconv"

// BinaryGap : returns the binary gap of an integer
func BinaryGap(N int) int {
	binary := strconv.FormatInt(int64(N), 2)
	binaryChars := []rune(binary)
	higher := 0
	acc := 0
	for i := 0; i < len(binaryChars); i++ {
		// 48 == "0"| 49 == "1"
		if binaryChars[i] == 49 {
			if acc > higher {
				higher = acc
			}
			acc = 0
		}
		if binaryChars[i] == 48 {
			acc++
		}
	}
	return higher
}
