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

	first := piscine.Atoi(os.Args[1])
	second := piscine.Atoi(os.Args[2])

	if first == second {
		piscine.PrintStr(piscine.Itoa(first))
		z01.PrintRune('\n')
		return
	}
	if piscine.IsPrime(first) {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}
	if piscine.IsPrime(second) {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	for first != second {
		if first > second {
			first = first - second
		} else {
			second = second - first
		}
	}

	piscine.PrintStr(piscine.Itoa(first))
	z01.PrintRune('\n')
}
