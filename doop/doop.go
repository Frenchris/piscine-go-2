package main

import (
	"fmt"
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {

	arr := os.Args

	lengthArgs := piscine.ArrayStrLength(arr)

	if lengthArgs != 4 {
		fmt.Println("a")
		return
	}

	if !piscine.IsNumeric(arr[1]) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	if !piscine.IsNumeric(arr[3]) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	if arr[2] == "+" {

		fmt.Print(piscine.Atoi(arr[1]) + piscine.Atoi(arr[3]))
		z01.PrintRune('\n')

	} else if arr[2] == "-" {

		fmt.Print(piscine.Atoi(arr[1]) - piscine.Atoi(arr[3]))
		z01.PrintRune('\n')

	} else if arr[2] == "*" {

		fmt.Print(piscine.Atoi(arr[1]) * piscine.Atoi(arr[3]))
		z01.PrintRune('\n')

	} else if arr[2] == "/" {

		if arr[3] == "0" {
			fmt.Print("No division by 0")
			z01.PrintRune('\n')
			return
		} else {
			fmt.Print(piscine.Atoi(arr[1]) / piscine.Atoi(arr[3]))
			z01.PrintRune('\n')
		}

	} else if arr[2] == "%" {

		if arr[3] == "0" {
			fmt.Print("No modulo by 0")
			z01.PrintRune('\n')
			return
		} else {
			fmt.Print(piscine.Atoi(arr[1]) % piscine.Atoi(arr[3]))
			z01.PrintRune('\n')
		}

	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
}
