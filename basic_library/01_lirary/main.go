package main

import (
	"fmt"

	"github.com/muhfaris/goFun/basic_library/01_lirary/arithmetic"
	"github.com/muhfaris/goFun/basic_library/01_lirary/devide"
	"github.com/muhfaris/goFun/basic_library/01_lirary/multiply"
	"github.com/muhfaris/goFun/basic_library/01_lirary/subtract"
)

func main() {
	sumRes := arithmetic.Sum(5, 6)
	subRes := subtract.Subtract(7, 3)
	multipleRes := multiply.Multiply(10, 4)
	devideRes, _ := devide.Devide(35, 7)

	fmt.Printf("5+6=%d, 7-3=%d, 10*4=%d, 35/7=%f \n", sumRes, subRes, multipleRes, devideRes)
}
