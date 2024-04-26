package main

import "errors"

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50 {
		return 0, errors.New("Скидка не может превышать 50%")
	}
	return price - price*discount/100, nil
}
