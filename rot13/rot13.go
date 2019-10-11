package main

import (
	"os"

	"github.com/01-edu/z01"

	piscine ".."
)

func main() {

	if piscine.ArrayStrLength(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	s := []rune(os.Args[1])

	for i := 0; i < piscine.StrLen(os.Args[1]); i++ {
		if (s[i] >= 'A' && s[i] <= 'L') ||
			(s[i] >= 'a' && s[i] <= 'l') {
			s[i] += 13
		} else if (s[i] > 'L' && s[i] <= 'Z') ||
			(s[i] > 'l' && s[i] <= 'z') {
			s[i] -= 13
		}
	}

	piscine.PrintStr(string(s))
	z01.PrintRune('\n')
}
