package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	noValid := ""
	negative := false

	if StrLen(base) == 10 {
		fmt.Print(nbr)
		return
	}

	if StrLen(base) < 2 {
		noValid = "NV"
	} else {
		for i := 0; i < StrLen(base)-1; i++ {
			for j := i + 1; j < StrLen(base); j++ {
				if base[i] == base[j] ||
					base[i] == '+' || base[i] == '-' ||
					base[j] == '+' || base[j] == '-' {
					noValid = "NV"
				}
			}
		}
	}

	if noValid == "NV" {
		PrintStr(noValid)
	} else {
		if nbr < 0 {
			nbr *= -1
			negative = true
		}
		SeeNumber(nbr, base, negative)
	}

}

func SeeNumber(nbr int, base string, negative bool) {

	array := []rune(base)

	if nbr >= 0 && nbr < StrLen(base) {
		if negative {
			z01.PrintRune('-')
			z01.PrintRune(array[nbr])
			return
		}
		z01.PrintRune(array[nbr])
		return

	}

	SeeNumber(nbr/StrLen(base), base, negative)
	z01.PrintRune(array[nbr%StrLen(base)])
}
