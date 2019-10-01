package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {

	//args 2-n

	fileRe, err := os.Open(os.Args[3])

	if err != nil {
		piscine.PrintStr(err.Error())
		z01.PrintRune('\n')
		return
	}

	stat, _ := fileRe.Stat()

	arr := make([]byte, stat.Size())

	fileRe.Read(arr)

	arrLength := piscine.StrLen(string(arr))

	piscine.PrintStr(string(arr[arrLength-piscine.Atoi(os.Args[2]) : arrLength]))

}
