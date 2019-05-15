package kata

// PlayPass transforms string
//
// your text: "BORN IN 2015!", shift 1
// 1 + 2 + 3 -> "CPSO JO 7984!"
// 4 "CpSo jO 7984!"
// 5 "!4897 Oj oSpC"
func PlayPass(s string, n int) string {
	// 1. shift each letter by a given number but the transformed letter must be a letter (circular shift),
	// 2. replace each digit by its complement to 9,
	// 3. keep such as non alphabetic and non digit characters,
	// 4. downcase each letter in odd position, upcase each letter in even position (the first character is in position 0),
	// 5. reverse the whole result.

	var r string

	return r
}
