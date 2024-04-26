package main

import (
	"time"
)

type Indicatorer interface {
	SMA(pair string, limit, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

type Indicator struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
	calculateEMA func(data []float64, period int) []float64
}

type IndicatorOption func(*Indicator)

func WithSMA(calculateSMA func(data []float64, period int) []float64) IndicatorOption {
	return func(i *Indicator) {
		i.calculateSMA = calculateSMA
	}
}

func WithEMA(calculateEMA func(data []float64, period int) []float64) IndicatorOption {
	return func(i *Indicator) {
		i.calculateEMA = calculateEMA
	}
}

func calculateSMA(data []float64, period int) []float64 {
	var sum float64
	res := make([]float64, 0, len(data)/period+1)
	i := 0
	for _, dd := range data {
		sum += dd
		i++
		if i == period {
			res = append(res, sum/float64(period))
			sum = 0
			i = 0
		}
	}
	return res
}

func calculateEMA(data []float64, period int) []float64 {
	alpha := 2.0 / (float64(period + 1))

	for i := 1; i < len(data); i++ {
		data[i] = alpha*data[i] + (1-alpha)*data[i-1]
	}
	return data
}

func NewIndicator(exchange Exchanger, opts ...IndicatorOption) *Indicator {
	i := &Indicator{
		exchange:     exchange,
		calculateEMA: calculateEMA,
		calculateSMA: calculateSMA,
	}
	for _, opt := range opts {
		opt(i)
	}
	return i
}

func (i *Indicator) SMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	prices, err := i.exchange.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return []float64{}, err
	}
	res := i.calculateSMA(prices, period)
	return res, err
}

func (i *Indicator) EMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	prices, err := i.exchange.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return []float64{}, err
	}
	res := i.calculateEMA(prices, period)
	return res, err
}

/*
func main() {
	var exchange Exchanger
	exchange = NewExmo(WithClient(&http.Client{}), WithURL("https://api.exmo.com/v1.1"))
	indicator := NewIndicator(exchange, WithEMA(calculateEMA), WithSMA(calculateSMA))

	sma, err := indicator.SMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sma)

	ema, err := indicator.EMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ema)
}
*/
