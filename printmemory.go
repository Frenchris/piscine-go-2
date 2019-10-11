package piscine

import (
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]int) {
	hexa := make([]string, 10)
	runes := make([]rune, 10)

	for i := 0; i < 10; i++ {
		runes[i] = rune(arr[i])
	}

	for i := 0; i < 10; i++ {
		hexa[i] = ConvertFromDecimal(arr[i], "0123456789abcdef")
	}
	for i := 0; i < 10; i++ {
		hexa[i] = EightLength(hexa[i])
	}
	for i := 0; i < 10; i++ {
		PrintStr(hexa[i])
		z01.PrintRune(' ')
		if i == 3 || i == 7 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')

	for i := 0; i < 10; i++ {
		if runes[i] < 31 {
			PrintStr(".")
		} else {
			PrintStr(string(runes[i]))
		}
	}
	z01.PrintRune('\n')
}

func EightLength(str string) string {
	for StrLen(str) != 9 {
		if StrLen(str) == 4 {
			str += " "
		}
		str += "0"
	}
	return str
}
