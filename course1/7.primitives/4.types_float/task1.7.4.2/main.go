package main

import (
	"time"

	"github.com/mattevans/dinero"
)

func currencyPairRate(from, to string, value float64) float64 {
	client := dinero.NewClient(
		"827dec1f583440849866926eb3f16106",
		from,
		20*time.Minute,
	)
	a, err := client.Rates.Get(to)
	if err != nil {
		return 0
	}
	return *a * 100
}
