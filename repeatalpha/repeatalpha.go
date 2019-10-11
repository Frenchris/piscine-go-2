package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	if piscine.ArrayStrLength(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	array := []rune(os.Args[1])
	var result []rune

	for i := 0; i < piscine.ArrayRuneLength(array); i++ {
		if array[i] >= 'a' && array[i] <= 'z' {
			for j := 'a'; j < 'z'; j++ {
				if array[i] == j {
					for k := 1; k <= int(j)-96; k++ {
						result = append(result, j)
					}
				}
			}
		} else if array[i] >= 'A' && array[i] <= 'Z' {
			for j := 'A'; j < 'Z'; j++ {
				if array[i] == j {
					for k := 1; k <= int(j)-64; k++ {
						result = append(result, j)
					}
				}
			}
		} else {
			result = append(result, array[i])
		}
	}

	piscine.PrintStr(string(result))
	z01.PrintRune('\n')
}
