package main

import (
	"os"

	piscine ".."

	"github.com/01-edu/z01"
)

func main() {

	for i := 1; i < piscine.ArrayStrLength(os.Args); i++ {
		piscine.PrintStr(os.Args[i])
		z01.PrintRune('\n')
	}

}
