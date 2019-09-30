package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for i := 0; i < ArrayStrLength(table); i++ {
		PrintStr(table[i])
		z01.PrintRune('\n')
	}
}
