package kata

import (
	"strings"
)

func reverseString(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

// IsPalindrome tests if string is a palindrome
func IsPalindrome(str string) bool {
	str = strings.ToUpper(str)

	return str == reverseString(str)
}
