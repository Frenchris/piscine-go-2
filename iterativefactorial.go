package piscine

func IterativeFactorial(nb int) int{
	

	num := 0

	if nb < 0 {
		return 0
	}

	if nb == 0{
		num = 1 
	}else{
		num = 1
		for i := 1; i <= nb; i++ {
			num *= i
		}
	}

	return num
}