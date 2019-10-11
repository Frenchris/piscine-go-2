package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	if piscine.ArrayStrLength(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	first := os.Args[1]
	second := os.Args[2]

	if first == "" {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}
	if second == "" {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	i := 0
	j := 0
	count := 0

	firstA := []rune(first)
	secondA := []rune(second)

	for i < piscine.StrLen(first) {
		for j < piscine.StrLen(second) {
			if firstA[i] == secondA[j] {
				count++
				i++
			}
			if i == piscine.StrLen(first) {
				break
			}
			j++
		}
		i++
	}

	if count == piscine.StrLen(first) {
		z01.PrintRune('1')

	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')

}
