package sample1

import (
	"fmt"
	"strconv"

	"github.com/huynvk/gosample/calc"
)

func Calc(input string) (float64, error) {
	a, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	b, err := calc.Sqrt(a)
	if err != nil {
		return 0, err
	}

	result, err := calc.Divide(1, b)
	return result, err
}

func Handle() {
	fmt.Println(Calc("25"))
	fmt.Println(Calc("-1"))
	fmt.Println(Calc("0"))
}
