package main

import (
	"os"

	"github.com/01-edu/z01"

	piscine ".."
)

func main() {
	if piscine.ArrayStrLength(os.Args) != 4 {
		z01.PrintRune('\n')
		return
	}

	var arr []rune
	var secondArg []rune
	var thirdArg []rune

	for i := 0; i < piscine.StrLen(os.Args[1]); i++ {
		arr = []rune(os.Args[1])
	}

	for i := 0; i < piscine.StrLen(os.Args[2]); i++ {
		secondArg = []rune(os.Args[2])

	}

	for i := 0; i < piscine.StrLen(os.Args[3]); i++ {
		thirdArg = []rune(os.Args[3])

	}

	for i := 0; i < piscine.StrLen(os.Args[1]); i++ {
		if arr[i] == secondArg[0] {
			arr[i] = thirdArg[0]
		}
	}

	result := ""

	for i := 0; i < piscine.StrLen(os.Args[1]); i++ {
		result += string(arr[i])
	}
	result += "\n"
	piscine.PrintStr(result)
}
