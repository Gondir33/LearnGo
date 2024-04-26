package main

import (
	"errors"

	"github.com/shopspring/decimal"
)

func DecimalSum(a, b string) (string, error) {
	x, err1 := decimal.NewFromString(a)
	y, err2 := decimal.NewFromString(b)
	if err1 != nil {
		return "0", err1
	} else if err2 != nil {
		return "0", err2
	} else {
		return decimal.Sum(x, y).String(), err1
	}
}

func DecimalSubtract(a, b string) (string, error) {
	x, err1 := decimal.NewFromString(a)
	y, err2 := decimal.NewFromString(b)
	if err1 != nil {
		return "0", err1
	} else if err2 != nil {
		return "0", err2
	} else {
		return x.Sub(y).String(), err1
	}
}
func DecimalMultiply(a, b string) (string, error) {
	x, err1 := decimal.NewFromString(a)
	y, err2 := decimal.NewFromString(b)
	if err1 != nil {
		return "0", err1
	} else if err2 != nil {
		return "0", err2
	} else {
		return x.Mul(y).String(), err1
	}
}

func DecimalDivide(a, b string) (string, error) {
	x, err1 := decimal.NewFromString(a)
	y, err2 := decimal.NewFromString(b)
	if err1 != nil {
		return "0", err1
	} else if err2 != nil {
		return "0", err2
	} else if y.IsZero() == true {
		return "0", errors.New("Divide to zero")
	} else {
		return x.Div(y).String(), err1
	}
}
func DecimalRound(a string, precision int32) (string, error) {
	x, err1 := decimal.NewFromString(a)
	if err1 != nil {
		return "0", err1
	}
	return x.Round(precision).String(), err1
}
func DecimalGreaterThan(a, b string) (bool, error) {
	x, err1 := decimal.NewFromString(a)
	y, err2 := decimal.NewFromString(b)
	if err1 != nil {
		return false, err1
	} else if err2 != nil {
		return false, err2
	} else {
		if x.Cmp(y) == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
}
func DecimalLessThan(a, b string) (bool, error) {
	x, err1 := decimal.NewFromString(a)
	y, err2 := decimal.NewFromString(b)
	if err1 != nil {
		return false, err1
	} else if err2 != nil {
		return false, err2
	} else {
		if x.Cmp(y) == -1 {
			return true, nil
		} else {
			return false, nil
		}
	}
}
func DecimalEqual(a, b string) (bool, error) {
	x, err1 := decimal.NewFromString(a)
	y, err2 := decimal.NewFromString(b)
	if err1 != nil {
		return false, err1
	} else if err2 != nil {
		return false, err2
	} else {
		if x.Cmp(y) == 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
}
