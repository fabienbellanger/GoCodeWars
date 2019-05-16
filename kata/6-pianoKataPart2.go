package kata

// WhichNote will receive an integer between 1 and 10000 (maybe you think that in principle it would be cool to count up to, say, a billion,
// but considering how many years it would take it is just not possible) and return one of the strings
// "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", or "G#" indicating which note you stopped on
// 1 <= keyPressCount <= 10000
func WhichNote(keyPressCount int) string {
	keysPattern := [12]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

	return keysPattern[((keyPressCount-1)%88)%12]
}
