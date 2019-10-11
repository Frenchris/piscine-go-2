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

	nbr := piscine.Atoi(os.Args[1])
	number := piscine.ConvertFromDecimal(nbr, "0123456789ABCDEF")
	number = piscine.LowerCase(number)

	piscine.PrintStr(number)
	z01.PrintRune('\n')
}
