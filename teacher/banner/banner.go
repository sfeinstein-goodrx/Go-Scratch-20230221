package main

import (
	"fmt"
	"unicode/utf8"
)

/*
banner("Go", 6)
  Go
------
*/

func main() {
	banner("Go", 6)
	s := "G☺"
	banner(s, 6)

	fmt.Println("len:", len("G☺"), "s[0]:", s[0]) // s[0] is a "byte"
	for i, r := range s {
		fmt.Println(i, r) // r is a "rune"
	}
	// s[0] = 'J' // strings are immutable
}

func banner(text string, width int) {
	offset := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
