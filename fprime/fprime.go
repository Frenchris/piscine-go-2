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
	count := 0
	l := append(PrimeFactors(piscine.Atoi(os.Args[1])))
	for i := 0; i < len(l); i++ {
		count++
		if count >= i && i > 0 {
			piscine.PrintStr("*")
		}
		piscine.PrintStr(piscine.Itoa(l[i]))

	}
	z01.PrintRune('\n')

}

func PrimeFactors(x int) []int {
	factors := []int{}

	candidate := 2
	for candidate <= x {
		if x%candidate == 0 {
			factors = append(factors, candidate)
			x = x / candidate
		} else {
			candidate++
		}
	}

	return factors
}
