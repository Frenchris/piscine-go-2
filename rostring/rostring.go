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

	array := piscine.Split(os.Args[1], " ")
	array = piscine.RemoveFromStrArray(array, "")

	if piscine.ArrayStrLength(array) == 1 {
		piscine.PrintStr(array[0])
		z01.PrintRune('\n')
		return
	}

	array = append(array, array[0])

	for i := 1; i < piscine.ArrayStrLength(array); i++ {
		piscine.PrintStr(array[i])
		z01.PrintRune(' ')
	}

	z01.PrintRune('\n')

}
