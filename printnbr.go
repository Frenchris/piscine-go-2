package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int){

	if n < 0 {
		z01.PrintRune('-')
		
		n = n*-1
	}
	if n != 0 {
		q := (n / 10)
		if q != 0{
			PrintNbr(q)
		}
		d := (n % 10) +'0'
		z01.PrintRune(rune(d))
	} else {
		z01.PrintRune('0')
	}	
}


