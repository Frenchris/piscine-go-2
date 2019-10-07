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

	firstArg := piscine.RemoveDuplicates(os.Args[1])
	secondArg := os.Args[2]

	piscine.PrintStr(piscine.CommonLetters(firstArg, secondArg))
}
