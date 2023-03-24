package main

import (
	"fmt"
)

func main() {
	i := 0
	runeText := [7]rune{1057, 1040, 1064, 1040, 1056, 1042, 1054}                              // create 7 Russian alphabet with rune (32 bit) in Decimal
	var text = [7]string{"\u0421", "\u0410", "\u0428", "\u0410", "\u0420", "\u0412", "\u041E"} // create 7 Russian alphabet with array of string

	for j := 0; j < 11; j++ {
		for i < 5 {
			fmt.Printf("Nilai i = %d\n", i)
			i++
		}

		if j == 5 {
			for b := 0; b < 14; b += 2 {
				letter := text[b/2] // get utf8 encoding string for current alphabet
				fmt.Printf("character %U %q starts at byte position %d\n", runeText[b/2], letter, b)
			}
		} else {
			fmt.Printf("Nilai j = %d\n", j)
		}

	}
}
