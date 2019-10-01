package main

import (
	"os"

	"github.com/01-edu/z01"

	piscine ".."
)

func main() {

	arguments := os.Args
	nbrArguments := piscine.ArrayStrLength(arguments)

	if nbrArguments >= 2 {
		for i := 1; i < nbrArguments; i++ {
			if os.Args[i] == "quest8.txt" {
				if i+1 == nbrArguments {
					piscine.PrintFile("quest8.txt")
				} else {
					piscine.PrintFile("quest8.txt")
					z01.PrintRune('\n')
				}
			}
			if os.Args[i] == "quest8T.txt" {
				if i+1 == nbrArguments {
					piscine.PrintFile("quest8T.txt")
				} else {
					piscine.PrintFile("quest8T.txt")
					z01.PrintRune('\n')
				}
			}
		}
	}
}
