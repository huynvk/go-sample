package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("Sqrt of negative number")
	}

	return math.Sqrt(number), nil
}

type DivisionError struct {
	Msg string
}

func (e *DivisionError) Error() string {
	return e.Msg
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivisionError{Msg: "Division by zero"}
	}

	return a / b, nil
}

func main() {
	fmt.Println(Divide(1, 0))
	fmt.Println(Divide(1, 2))
	fmt.Println(Sqrt(-1))
	fmt.Println(Sqrt(25))
}
