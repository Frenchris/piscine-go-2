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
			fileRe, err := os.Open(os.Args[i])

			if err != nil {
				piscine.PrintStr(err.Error())
				z01.PrintRune('\n')

				return
			}
			stat, _ := fileRe.Stat()

			arr := make([]byte, stat.Size())

			fileRe.Read(arr)

			if i+1 == nbrArguments {
				piscine.PrintStr(string(arr))
				z01.PrintRune('\n')
			} else {
				piscine.PrintStr(string(arr))
				z01.PrintRune('\n')
				z01.PrintRune('\n')
			}

		}
	}
}
