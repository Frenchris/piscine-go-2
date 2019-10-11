package main

import (
	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	str := "abcdefghijklmnopqrstuvwxyz"
	arr := []rune(str)
	for i := 1; i < piscine.StrLen(str); i++ {
		if i%2 != 0 {
			arr[i] -= 32
		}
	}
	for i := 0; i < piscine.ArrayRuneLength(arr); i++ {
		z01.PrintRune(arr[i])
	}
	z01.PrintRune('\n')
}
