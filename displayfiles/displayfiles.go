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
	} else {

		fileWr, err := os.Create("quest8.txt")
		if err != nil {
			piscine.PrintStr(err.Error())
			return
		}

		fileWr.WriteString("Almost there!!")

		fileRe, err := os.Open("quest8.txt")
		if err != nil {
			piscine.PrintStr(err.Error())
		}

		var arr []byte

		fileRe.Read(arr)

		piscine.PrintStr(string(arr))
		fileRe.Close()

		z01.PrintRune('\n')

	}

}
