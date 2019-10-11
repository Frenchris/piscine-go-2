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

	for i := 1; i <= 9; i++ {
		piscine.PrintStr(piscine.Itoa(i))
		piscine.PrintStr(" x ")
		piscine.PrintStr(os.Args[1])
		piscine.PrintStr(" = ")
		result := i * piscine.Atoi(os.Args[1])

		piscine.PrintStr(piscine.Itoa(result))
		z01.PrintRune('\n')
	}
}
