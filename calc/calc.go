package calc

import (
	"errors"
	"math"
)

func Divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("Division by zero")
	}

	return dividend / divisor, nil
}

func Sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("Sqrt of negative number")
	}

	return math.Sqrt(number), nil
}
