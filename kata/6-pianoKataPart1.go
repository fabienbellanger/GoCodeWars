package kata

// BlackOrWhiteKey will receive an integer between 1 and 10000 and return the string "black" or "white"
// 1 <= keyPressCount <= 10000
func BlackOrWhiteKey(keyPressCount int) string {
	keysPattern := [12]string{"white", "black", "white", "white", "black", "white", "black", "white", "white", "black", "white", "black"}

	return keysPattern[((keyPressCount-1)%88)%12]
}
