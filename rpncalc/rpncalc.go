package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

var result int

func main() {

	if piscine.ArrayStrLength(os.Args) != 2 {
		piscine.PrintStr("Error")
		z01.PrintRune('\n')
		return
	}

	args := piscine.Split(os.Args[1], " ")
	args = piscine.RemoveFromStrArray(args, "")

	if !piscine.CheckStrings(args) {
		piscine.PrintStr("Error")
		z01.PrintRune('\n')
		return
	}

	operands := 0
	operators := 0

	for i := 0; i < piscine.ArrayStrLength(args); i++ {
		if piscine.IsNumber(args[i]) {
			operands++
		}
		if piscine.IsOperator(args[i]) {
			operators++
		}
	}

	i := 0
	j := i + 1
	k := j + 1

	result := 0

	if operands-1 == operators {
		for i < piscine.ArrayStrLength(args)-2 {
			for j < piscine.ArrayStrLength(args)-1 {
				for k < piscine.ArrayStrLength(args) {
					if piscine.ArrayStrLength(args) >= 3 {
						if piscine.IsNumber(args[i]) && piscine.IsNumber(args[j]) && piscine.IsOperator(args[k]) {
							switch args[k] {
							case "+":
								result = piscine.Atoi(args[i]) + piscine.Atoi(args[j])
								args[k] = piscine.Itoa(piscine.Atoi(args[i]) + piscine.Atoi(args[j]))
							case "-":
								result = piscine.Atoi(args[i]) - piscine.Atoi(args[j])
								args[k] = piscine.Itoa(piscine.Atoi(args[i]) - piscine.Atoi(args[j]))
							case "*":
								result = piscine.Atoi(args[i]) * piscine.Atoi(args[j])
								args[k] = piscine.Itoa(piscine.Atoi(args[i]) * piscine.Atoi(args[j]))
							case "/":
								result = piscine.Atoi(args[i]) / piscine.Atoi(args[j])
								args[k] = piscine.Itoa(piscine.Atoi(args[i]) / piscine.Atoi(args[j]))
							case "%":
								result = piscine.Atoi(args[i]) % piscine.Atoi(args[j])
								args[k] = piscine.Itoa(piscine.Atoi(args[i]) % piscine.Atoi(args[j]))
							}
							args = args[2:piscine.ArrayStrLength(args)]
						} else {
							i++
							j++
							k++
						}
					}
				}
			}
		}
	} else {
		piscine.PrintStr("Error")
		z01.PrintRune('\n')
		return
	}

	piscine.PrintStr(piscine.Itoa(result))
	z01.PrintRune('\n')
}
