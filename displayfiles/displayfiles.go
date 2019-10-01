package main

import (
	"io/ioutil"
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

		file, err := ioutil.ReadFile("quest8.txt")
		if err != nil {
			piscine.PrintStr(err.Error())
			return
		}

		piscine.PrintStr(string(file))
		z01.PrintRune('\n')

	}

}
