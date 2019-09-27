package printparams

import (
	"os"

	piscine ".."
)

func main() {

	for i := 1; i < piscine.ArrayStrLength(os.Args); i++ {
		piscine.PrintStr(os.Args[i])
	}

}
