package main

import "fmt"
import "github.com/fabienbellanger/GoCodeWars/kata"

func main() {
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

	s = "BORN IN 2015!"
	n = 1
	fmt.Printf("s = %s ; n = %d => <%s> | <!4897 Oj oSpC>\n", s, n, kata.PlayPass(s, n))
}
