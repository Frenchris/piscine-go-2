package main

import (
	"os"

	piscine ".."
)

func main() {

	if piscine.ArrayStrLength(os.Args) == 1 {
		piscine.PrintStr("\n")
		return
	}

	arr := make([]byte, 2048)
	p := &arr[0]
	arg := []rune(os.Args[1])

	i := 0
	j := 0

	for j < 2048 {
		for i < piscine.ArrayRuneLength(arg) {
			switch arg[i] {
			case '>':
				j++
				p = &arr[j]
			case '<':
				j--
				p = &arr[j]
			case '+':
				*p++
			case '-':
				*p--
			case '.':
				piscine.PrintStr(string(*p))
			case '[':
				cont := 0

				if *p == 0 {

					for cont >= 0 {
						i++
						if arg[i] == ']' {
							cont--
						} else if arg[i] == '[' {
							cont++
						}
					}
				}
			case ']':
				cont := 0
				if *p != 0 {
					for cont >= 0 {
						i--
						if arg[i] == '[' {
							cont--
						} else if arg[i] == ']' {
							cont++
						}
					}
				}
			}

			i++
		}
		j = 2048
	}
}
