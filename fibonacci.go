package piscine

func Fibonacci(index int) int{
	if index <= 0{
		return 0
	} 
	if index == 1{
		return 1
	} else {
		return Fibonacci(index - 1) + Fibonacci(index - 2)
	}
}