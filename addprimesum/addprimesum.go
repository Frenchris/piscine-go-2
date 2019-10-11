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
	result := 0
	primes := piscine.PrimeNumbers(piscine.Atoi(os.Args[1]))

	for i := 0; i < piscine.ArrayIntLength(primes); i++ {
		result += primes[i]
	}

	piscine.PrintStr(piscine.Itoa(result))
	z01.PrintRune('\n')
}
