package main

import (
	"time"
)

type GeneralIndicatorer interface {
	GetData(pair string, limit, period int, from, to time.Time, indicator Indicatorer) ([]float64, error)
}

type GeneralIndicator struct{}

func (gi *GeneralIndicator) GetData(pair string, limit, period int, from, to time.Time, indicator Indicatorer) ([]float64, error) {
	return indicator.GetData(pair, limit, period, from, to)
}

type Indicatorer interface {
	GetData(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

type IndicatorSMA struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
}

type IndicatorEMA struct {
	exchange     Exchanger
	calculateEMA func(data []float64, period int) []float64
}

func NewIndicatorSMA(exchange Exchanger) *IndicatorSMA {
	is := &IndicatorSMA{exchange: exchange}
	is.calculateSMA = calculateSMA
	return is
}

func NewIndicatorEMA(exchange Exchanger) *IndicatorEMA {
	is := &IndicatorEMA{exchange: exchange}
	is.calculateEMA = calculateEMA
	return is
}

func (i *IndicatorEMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	prices, err := i.exchange.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return []float64{}, err
	}
	return i.calculateEMA(prices, period), nil
}

func (i *IndicatorSMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	prices, err := i.exchange.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return []float64{}, err
	}
	return i.calculateSMA(prices, period), nil
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

/*
func main() {
	var exchange Exchanger
	exchange = NewExmo(WithClient(&http.Client{}), WithURL("https://api.exmo.com/v1.1"))
	indicatorSMA := NewIndicatorSMA(exchange)
	generalIndicator := &GeneralIndicator{}
	sma, err := generalIndicator.GetData("BTC_USD", 120, 10, time.Now().AddDate(0, 0, -2), time.Now(), indicatorSMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sma)

	indicatorEMA := NewIndicatorEMA(exchange)
	ema, err := generalIndicator.GetData("BTC_USD", 120, 10, time.Now().AddDate(0, 0, -2), time.Now(), indicatorEMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ema)
}
*/
