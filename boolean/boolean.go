package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {

	if nbr%2 == 0 {
		return true
	}
	return false

}

func main() {
	lengthOfArg := piscine.ArrayStrLength(os.Args) - 1
	if isEven(lengthOfArg) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
