package sample2

import (
	"fmt"
	"strconv"

	"github.com/huynvk/gosample/calc"
)

type any interface{}

type pair struct {
	result float64
	err    error
}

func newPair(result float64, err error) *pair {
	return &pair{result: result, err: err}
}

func (p *pair) execute(wrapper func(float64) (float64, error)) *pair {
	if p.err != nil {
		return p
	}
	result, err := wrapper(p.result)
	return newPair(result, err)
}

func divide(input float64) (float64, error) {
	return calc.Divide(1, input)
}

func Calc(input string) (float64, error) {
	a := newPair(strconv.ParseFloat(input, 64)).execute(calc.Sqrt).execute(divide)
	return float64(a.result), a.err
}

func Handle() {
	fmt.Println(Calc("25"))
	fmt.Println(Calc("-1"))
	fmt.Println(Calc("0"))
}
