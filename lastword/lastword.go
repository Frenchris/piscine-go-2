package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {

	if piscine.ArrayStrLength(os.Args) != 2 ||
		piscine.CountInstances(os.Args[1], " ") == piscine.StrLen(os.Args[1]) {
		z01.PrintRune('\n')
		return
	}
	array := piscine.Split(os.Args[1], " ")
	last := array[piscine.ArrayStrLength(array)-1]

	for last == " " || last == "" {
		array = array[0 : piscine.ArrayStrLength(array)-1]
		last = array[piscine.ArrayStrLength(array)-1]
	}

	piscine.PrintStr(last)
	z01.PrintRune('\n')
}
