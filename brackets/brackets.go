package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	if piscine.ArrayStrLength(os.Args) == 1 {
		z01.PrintRune('\n')
		return
	}

	for i := 1; i < piscine.ArrayStrLength(os.Args); i++ {
		arr := []rune(os.Args[i])
		var stack []rune

		if os.Args[i] == "" {
			piscine.PrintStr("OK")
			z01.PrintRune('\n')
		} else {
			for j := 0; j < piscine.StrLen(os.Args[i]); j++ {
				if arr[j] == '(' || arr[j] == '[' || arr[j] == '{' {
					stack = append(stack, arr[j])
				} else {
					switch arr[j] {
					case ')':
						if stack[piscine.ArrayRuneLength(stack)-1] == '(' {
							stack = stack[0 : piscine.ArrayRuneLength(stack)-1]

						}
					case ']':
						if stack[piscine.ArrayRuneLength(stack)-1] == '[' {
							stack = stack[0 : piscine.ArrayRuneLength(stack)-1]

						}
					case '}':
						if stack[piscine.ArrayRuneLength(stack)-1] == '{' {
							stack = stack[0 : piscine.ArrayRuneLength(stack)-1]
						}
					}
				}
			}

			if piscine.ArrayRuneLength(stack) == 0 {
				piscine.PrintStr("OK")
				z01.PrintRune('\n')
			} else {
				piscine.PrintStr("Error")
				z01.PrintRune('\n')
			}

		}
	}
}
