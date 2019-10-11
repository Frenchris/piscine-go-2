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

	if !piscine.IsNumber(os.Args[1]) {
		piscine.PrintStr("Should be an int")
		z01.PrintRune('\n')
		return
	}

	nbr := piscine.Atoi(os.Args[1])

	for i := 0; i <= nbr; i++ {
		if nbr == piscine.Power(2, i) {
			piscine.PrintStr("true")
			z01.PrintRune('\n')
			return
		}
	}
	piscine.PrintStr("false")
	z01.PrintRune('\n')

}
