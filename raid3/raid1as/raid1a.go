package main

import (
	"os"

	piscine "../.."
)

func main() {

	piscine.Raid1a(piscine.Atoi(os.Args[1]), piscine.Atoi(os.Args[2]))
}
