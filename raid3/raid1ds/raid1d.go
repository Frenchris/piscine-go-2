package main

import (
	"os"

	piscine "../.."
)

func main() {

	piscine.Raid1d(piscine.Atoi(os.Args[1]), piscine.Atoi(os.Args[2]))
}
