package main

import "fmt"
import "github.com/fabienbellanger/GoCodeWars/kata"

func main() {
	var s string
	var n int

	s = "I LOVE YOU!!!"
	n = 1
	fmt.Printf("n = %d => %s | !!!vPz fWpM J", n, kata.PlayPass(s, n))

	n = 0
	fmt.Printf("n = %d => %s | !!!uOy eVoL I", n, kata.PlayPass(s, n))

	s = "AAABBCCY"
	n = 1
	fmt.Printf("n = %d => %s | zDdCcBbB", n, kata.PlayPass(s, n))
}
