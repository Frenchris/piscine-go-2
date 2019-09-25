package piscine

func IterativeFactorial(nb int) int{
	

	num := 0

	if nb < 0 {
		return num
	}
	num = 1
	if nb == 0{
		return num 
	}else{
		for i := 1; i <= nb; i++ {
			num *= i
		}
		return num
	}

	
}