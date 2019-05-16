package main

import "fmt"
import "github.com/fabienbellanger/GoCodeWars/kata"

func main() {
	// 8 kyu
	// -----
	// testPalindrome()
	// testTotalAmountOfPoint()

	// 7 kyu
	// -----
	// testDeodorantEvaporator()

	// 6 kyu
	// -----
	// testPassPhrases()
	// testPianoKataPart1()
	testPianoKataPart2()
}

func testPassPhrases() {
	fmt.Println("\n6 - PassPhrases")
	fmt.Printf("---------------\n")

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
	fmt.Println("\n8 - Palindrome")
	fmt.Printf("--------------\n")

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

func testTotalAmountOfPoint() {
	fmt.Println("\n8 - TotalAmountOfPoint")
	fmt.Printf("----------------------\n")

	m1 := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}
	fmt.Printf("%d | %d\n", kata.Points(m1), 30)

	m1 = []string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}
	fmt.Printf("%d | %d\n", kata.Points(m1), 0)
}

func testDeodorantEvaporator() {
	fmt.Println("\n7 - TotalAmountOfPoint")
	fmt.Printf("----------------------\n")

	c := 10.0
	e := 10
	t := 10
	fmt.Printf("%d | %d\n", kata.Evaporator(c, e, t), 22)

	c = 10.0
	e = 10
	t = 5
	fmt.Printf("%d | %d\n", kata.Evaporator(c, e, t), 29)

	c = 100.0
	e = 5
	t = 5
	fmt.Printf("%d | %d\n", kata.Evaporator(c, e, t), 59)
}

func testPianoKataPart1() {
	fmt.Println("\n6 - PianoKataPart1")
	fmt.Printf("------------------\n")

	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(1), "white")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(5), "black")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(12), "black")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(42), "white")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(88), "white")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(89), "white")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(92), "white")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(100), "black")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(111), "white")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(200), "black")
	fmt.Printf("%s | %s\n", kata.BlackOrWhiteKey(2017), "white")
}

func testPianoKataPart2() {
	fmt.Println("\n6 - PianoKataPart2")
	fmt.Printf("------------------\n")

	fmt.Printf("%s | %s\n", kata.WhichNote(1), "A")
	fmt.Printf("%s | %s\n", kata.WhichNote(5), "C#")
	fmt.Printf("%s | %s\n", kata.WhichNote(12), "G#")
	fmt.Printf("%s | %s\n", kata.WhichNote(42), "D")
	fmt.Printf("%s | %s\n", kata.WhichNote(88), "C")
	fmt.Printf("%s | %s\n", kata.WhichNote(89), "A")
	fmt.Printf("%s | %s\n", kata.WhichNote(92), "C")
	fmt.Printf("%s | %s\n", kata.WhichNote(100), "G#")
	fmt.Printf("%s | %s\n", kata.WhichNote(111), "G")
	fmt.Printf("%s | %s\n", kata.WhichNote(200), "G#")
	fmt.Printf("%s | %s\n", kata.WhichNote(2017), "F")
}
