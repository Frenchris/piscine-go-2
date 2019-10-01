package piscine

import (
	"github.com/01-edu/z01"
)

func ForEach(f func(int), arr []int) {
	for i := 0; i < ArrayIntLength(arr); i++ {
		f(arr[i])
	}
	z01.PrintRune('\n')
}
