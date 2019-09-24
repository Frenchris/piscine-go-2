package piscine

import (
	"github.com/01-edu/z01"
)

func Raid1a(x,y int){
	
	z01.PrintRune('o')

	for i := 0; i < x - 2; i++ {
		z01.PrintRune('-')
	}

	z01.PrintRune('o')
	z01.PrintRune('\n')

	for i := 0; i < y - 2; i++ {
		z01.PrintRune('|')
		for j := 0; j < x - 2; j++ {
			z01.PrintRune(' ')
		}
		z01.PrintRune('|')
		z01.PrintRune('\n')
	}
	
	z01.PrintRune('o')

	for i := 0; i < x - 2; i++ {
		z01.PrintRune('-')
	}

	z01.PrintRune('o')
	z01.PrintRune('\n')
}