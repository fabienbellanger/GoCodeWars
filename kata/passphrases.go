package kata

import (
	"fmt"
)

func find(a []rune, x rune) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(r []rune) []rune {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

// PlayPass transforms string
func PlayPass(s string, n int) string {
	// 1. shift each letter by a given number but the transformed letter must be a letter (circular shift),
	// 2. replace each digit by its complement to 9,
	// 3. keep such as non alphabetic and non digit characters,
	// 4. downcase each letter in odd position, upcase each letter in even position (the first character is in position 0),
	// 5. reverse the whole result.

	runes := make([]rune, 0)

	// A - Z : 65 - 90
	// a - z : 97 - 122
	// 0 - 9 : 48 - 57

	// Step 0
	// ------
	lettersRotation := make([]rune, 0)
	var i rune
	for i = 65; i <= 90; i++ {
		lettersRotation = append(lettersRotation, i)
	}
	for i = 97; i <= 122; i++ {
		lettersRotation = append(lettersRotation, i)
	}
	nbLetters := len(lettersRotation)
	fmt.Println(nbLetters)

	// Steps 1 - 2 - 3 - 4
	// -------------------
	for index, c := range s {
		if c >= 65 && c <= 90 || c >= 97 && c <= 122 {
			// Letter
			letterIndex := find(lettersRotation, c)
			newRune := lettersRotation[(letterIndex+n)%nbLetters]

			if index%2 == 1 {
				// Odd -> lowerCase
				if c <= 90 {
					newRune += 32
				}
			} else {
				// Even -> upperCase
				if c >= 97 {
					newRune -= 32
				}
			}
			runes = append(runes, newRune)
		} else if c >= 48 && c <= 57 {
			// Number
			newC := 57 - (c - 48)
			runes = append(runes, newC)
		} else {
			// Non alphabetic and non digit characters
			runes = append(runes, c)
		}
	}

	// Step 5
	// ------
	reverseRunes := Reverse(runes)

	return string(reverseRunes)
}
