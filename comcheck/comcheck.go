package main

import (
	"os"

	"github.com/01-edu/z01"

	piscine ".."
)

func main() {

	for i := 0; i < piscine.ArrayStrLength(os.Args); i++ {
		if os.Args[i] == "01" || os.Args[i] == "galaxy" || os.Args[i] == "galaxy 01" {
			piscine.PrintStr("Alert!!!")
			z01.PrintRune('\n')
			return
		}
	}
}
