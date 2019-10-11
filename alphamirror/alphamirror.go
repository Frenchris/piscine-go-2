package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {

	if piscine.ArrayStrLength(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	s := []rune(os.Args[1])

	for i := 0; i < piscine.StrLen(os.Args[1]); i++ {
		if s[i] >= 'a' && s[i] <= 'm' ||
			s[i] >= 'A' && s[i] <= 'M' {
			for j := 'a'; j <= 'm'; j++ {
				if s[i] == j {
					s[i] += 25 - ((j - 'a') * 2)
				}
			}
			for j := 'A'; j <= 'M'; j++ {
				if s[i] == j {
					s[i] += 25 - ((j - 'A') * 2)
				}
			}
		} else if s[i] >= 'n' && s[i] <= 'z' ||
			s[i] >= 'N' && s[i] <= 'Z' {
			for j := 'n'; j <= 'z'; j++ {
				if s[i] == j {
					s[i] = s[i] - 1 - ((j - 'n') * 2)
				}
			}
			for j := 'N'; j <= 'Z'; j++ {
				if s[i] == j {
					s[i] = s[i] - 1 - ((j - 'N') * 2)
				}
			}
		}
	}

	piscine.PrintStr(string(s))
	z01.PrintRune('\n')
}
