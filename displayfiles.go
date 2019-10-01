package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func DisplayFiles() {

	lengthArgs := ArrayStrLength(os.Args)
	if lengthArgs >= 3 {
		PrintStr("Too many arguments")
		z01.PrintRune('\n')
	} else if lengthArgs == 1 {
		PrintStr("File name missing")
		z01.PrintRune('\n')
	} else {
		PrintStr("Almost there!!")
	}

}
