package subtract

func Subtract(args ...int) int {
	if len(args) < 2 {
		return 0
	}
	res := args[0]
	for i := 1; i < len(args); i++ {
		res -= args[1]
	}
	return res
}
