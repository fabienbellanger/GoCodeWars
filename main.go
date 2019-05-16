package main

import "fmt"
import "github.com/fabienbellanger/GoCodeWars/kata"

func main() {
	// 8 kyu
	// -----
	// testPalindrome()

	// 6 kyu
	// -----
	// testPassPhrases()
}

func testPassPhrases() {
	var s string
	var n int

	s = "I LOVE YOU!!!"
	n = 1
	fmt.Printf("s = %s ; n = %d => <%s> | <!!!vPz fWpM J>\n", s, n, kata.PlayPass(s, n))

	n = 0
	fmt.Printf("s = %s ; n = %d => <%s> | <!!!uOy eVoL I>\n", s, n, kata.PlayPass(s, n))

	s = "AAABBCCY"
	n = 1
	fmt.Printf("s = %s ; n = %d => <%s> | <zDdCcBbB>\n", s, n, kata.PlayPass(s, n))

	s = "Z  BORN IN 2015!"
	n = 20
	fmt.Printf("s = %s ; n = %d => <%s> | <4897 hC HlIv  T>\n", s, n, kata.PlayPass(s, n))
}

func testPalindrome() {
	var s string
	var b bool

	s = "a"
	b = kata.IsPalindrome(s)
	fmt.Printf("%s | %t => %t\n", s, b, true)

	s = "aba"
	b = kata.IsPalindrome(s)
	fmt.Printf("%s | %t => %t\n", s, b, true)

	s = "abba"
	b = kata.IsPalindrome(s)
	fmt.Printf("%s | %t => %t\n", s, b, true)

	s = "hello"
	b = kata.IsPalindrome(s)
	fmt.Printf("%s | %t => %t\n", s, b, false)

	s = "aBbA"
	b = kata.IsPalindrome(s)
	fmt.Printf("%s | %t => %t\n", s, b, true)
}
