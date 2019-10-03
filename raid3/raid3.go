package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/01-edu/z01"

	piscine ".."
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output = append(output, input)
	}

	array := piscine.Split(string(output), "\n")
	if piscine.StrLen(string(output)) == 2 {
		if output[0] == 'o' {
			piscine.PrintStr("[raid1a] [1] [1]")
		} else if output[0] == '/' {
			piscine.PrintStr("[raid1b] [1] [1]")
		} else if output[0] == 'A' {
			piscine.PrintStr("[raid1c] [1] [1] || [raid1d] [1] [1] || [raid1e] [1] [1]")
		}
	} else {
		findParameters(array)
	}

	z01.PrintRune('\n')
}

func findParameters(arr []string) {

	x := piscine.StrLen(arr[0])
	y := piscine.ArrayStrLength(arr) - 1

	first := arr[0]
	last := arr[y-1]

	if first[0] == 'o' {
		piscine.PrintStr("[raid1a]")
	} else if first[0] == '/' {
		piscine.PrintStr("[raid1b]")
	} else if first[0] == 'A' {
		if first[x-1] == 'A' {
			piscine.PrintStr("[raid1c]")
		} else if last[0] == 'A' {
			piscine.PrintStr("[raid1d]")
		} else if last[0] == 'C' {
			piscine.PrintStr("[raid1e]")
		}

	}
	fmt.Printf(" [%v] [%v]", x, y)
}
