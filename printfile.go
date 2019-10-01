package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintFile(file string) {

	fileRe, err := os.Open(file)

	if err != nil {
		PrintStr(err.Error())
		z01.PrintRune('\n')

		return
	}
	stat, _ := fileRe.Stat()

	arr := make([]byte, stat.Size())

	fileRe.Read(arr)

	PrintStr(string(arr))
	z01.PrintRune('\n')

}
