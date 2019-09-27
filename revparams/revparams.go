package main

import (
	"os"

	piscine ".."

	"github.com/01-edu/z01"
)

func main() {

	for i := piscine.ArrayStrLength(os.Args) - 1; i > 0; i-- {
		piscine.PrintStr(os.Args[i])
		z01.PrintRune('\n')
	}

}
