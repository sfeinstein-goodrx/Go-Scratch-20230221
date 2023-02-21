package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G❤️", 6)
	banner("G☺", 6)
	report("apple", 6)
	report("banana", 23)
	report("G❤️", 23)
}
func report(item string, amount int) {
	fmt.Printf("%-10s %2d\n", item, amount)
}

func banner(text string, width int) {
	//offset := (width - len(text)) / 2
	runesCount := utf8.RuneCountInString(text)
	offset := (width - runesCount) / 2
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	fmt.Println("Runes in string", runesCount)
	for i, r := range text {
		fmt.Println(i, r) //r is a "rune"
	}
}
