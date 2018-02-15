package multiply

func Multiply(args ...int) int {
	if len(args) < 2 {
		return 0
	}

	res := 1
	for i := 0; i < len(args); i++ {
		res *= args[i]
	}
	return res
}
