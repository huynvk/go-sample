package sample3

import (
	"fmt"
	"strconv"

	"github.com/huynvk/gosample/calc"
)

func Calc(input string) float64 {
	a, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err.Error())
	}

	b, err := calc.Sqrt(a)
	if err != nil {
		panic(err.Error())
	}

	result, err := calc.Divide(1, b)
	if err != nil {
		panic(err.Error())
	}
	return result
}

func Handle() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(Calc("25"))
	fmt.Println(Calc("-1"))
	fmt.Println(Calc("0"))
}
