package devide

import "errors"

func Devide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("You can't devide by zero")
	}
	return float64(a) / float64(b), nil
}
