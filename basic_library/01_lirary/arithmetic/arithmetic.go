package arithmetic

func Sum(args ...int) (res int) {
	for _, v := range args {
		res += v
	}
	return
}
