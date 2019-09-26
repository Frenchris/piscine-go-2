package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintCombN(n int) {

	PrintCombina(0, n, "")
	z01.PrintRune('\n')

}
func PrintCombina(start, n int, out string) {

	// If number becomes N-digit, print it
	if n == 0 {
		if out == "9" || out == "89" || out == "789" ||
			out == "6789" || out == "56789" || out == "456789" ||
			out == "3456789" || out == "23456789" ||
			out == "123456789" {
			fmt.Print(out)
		} else {

			fmt.Print(out)
			z01.PrintRune(',')
			z01.PrintRune(' ')
			return
		}
	}

	// start from (prev digit + 1) till 9
	for i := start; i <= 9; i++ {

		number := ""

		switch i {
		case 0:
			number = "0"
		case 1:
			number = "1"
		case 2:
			number = "2"
		case 3:
			number = "3"
		case 4:
			number = "4"
		case 5:
			number = "5"
		case 6:
			number = "6"
		case 7:
			number = "7"
		case 8:
			number = "8"
		case 9:
			number = "9"
		}

		str := out + number

		// recurse for next digit
		PrintCombina(i+1, n-1, str)
	}
}

func PrintString(str string) {

	aStringChangeable := []rune(str)
	for i := 0; i < ArrayStrLength(str); i++ {

		//if aStringChangeable[i] >= ' ' && aStringChangeable[i] <= '~' {
		z01.PrintRune(aStringChangeable[i])
		//}
	}

}

func ArrayStrLength(str string) int {

	length := 0

	for range str {
		length++
	}

	return length
}
