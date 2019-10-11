package main

import (
	"os"

	"github.com/01-edu/z01"

	piscine ".."
)

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"

	if piscine.ArrayStrLength(os.Args) == 1 {
		piscine.PrintStr("options: " + alpha + "\n")
		return
	}

	alpharev := piscine.StrRev(alpha)

	numbers := "00000000 00000000 00000000 00000000"
	nbrs := []rune(numbers)

	var flags []rune

	for i := 1; i < piscine.ArrayStrLength(os.Args); i++ {
		flag := []rune(os.Args[i])
		if flag[0] == '-' {
			for j := 1; j < piscine.ArrayRuneLength(flag); j++ {
				if flag[j] >= 'a' && flag[j] <= 'z' {
					flags = append(flags, flag[j])
				}
			}
		} else {
			piscine.PrintStr("Invalid Option\n")
			return
		}
	}

	for i := 0; i < piscine.ArrayRuneLength(flags); i++ {
		helper := piscine.Index(alpharev, string(flags[i]))
		if flags[i] == 'z' || flags[i] == 'y' {
			helper += 6
		} else if flags[i] <= 'x' && flags[i] >= 'q' {
			helper += 7
		} else if flags[i] <= 'p' && flags[i] >= 'i' {
			helper += 8
		} else if flags[i] <= 'h' && flags[i] >= 'a' {
			helper += 9
		}
		nbrs[helper] = '1'
	}

	if nbrs[27] == '1' {
		piscine.PrintStr("options: " + alpha + "\n")
		return
	}
	for i := 0; i < piscine.ArrayRuneLength(nbrs); i++ {
		z01.PrintRune(nbrs[i])
	}
	z01.PrintRune('\n')
}
