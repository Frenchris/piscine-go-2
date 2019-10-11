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

	for i := 0; i < piscine.ArrayStrLength(array)/2; i++ {
		array[i], array[piscine.ArrayStrLength(array)-1-i] =
			array[piscine.ArrayStrLength(array)-1-i], array[i]
	}

	for j := 0; j < piscine.ArrayStrLength(array); j++ {
		piscine.PrintStr(array[j])
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
