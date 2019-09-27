package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintProgramName() {
	PrintStr(os.Args[0])
	z01.PrintRune('\n')
}
