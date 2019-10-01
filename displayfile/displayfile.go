package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {

	lengthArgs := piscine.ArrayStrLength(os.Args)
	if lengthArgs >= 3 {
		piscine.PrintStr("Too many arguments")
		z01.PrintRune('\n')
	} else if lengthArgs == 1 {
		piscine.PrintStr("File name missing")
		z01.PrintRune('\n')
	} else if os.Args[1] == "quest8.txt" {

		fileRe, err := os.Open("quest8.txt")

		if err != nil {
			piscine.PrintStr(err.Error())
			return
		}
		stat, _ := fileRe.Stat()

		arr := make([]byte, stat.Size())

		fileRe.Read(arr)

		piscine.PrintStr(string(arr))
		z01.PrintRune('\n')

	}

}
