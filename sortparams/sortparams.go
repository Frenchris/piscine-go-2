package main

import (
	"fmt"
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {

	fmt.Println(piscine.ArrayStrLength(os.Args))
	for i := 1; i < piscine.ArrayStrLength(os.Args)-1; i++ {
		for j := i + 1; j < piscine.ArrayStrLength(os.Args); j++ {
			if os.Args[i] > os.Args[j] {
				os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
			}
		}
	}

	for i := 1; i < piscine.ArrayStrLength(os.Args); i++ {
		piscine.PrintStr(os.Args[i])
		z01.PrintRune('\n')
	}

}
