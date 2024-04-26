package main

import "time"

type IndicatorWithCache struct {
	indicator Indicatorer
	cache     map[string][]float64
}

func (iwc *IndicatorWithCache) SMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	sma, ok := iwc.cache[pair]
	if ok {
		return sma, nil
	}
	sma, err := iwc.indicator.SMA(pair, limit, period, from, to)
	if err != nil {
		return []float64{}, err
	}
	iwc.cache[pair] = sma
	return sma, nil
}
func (iwc *IndicatorWithCache) EMA(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	ema, ok := iwc.cache[pair]
	if ok == true {
		return ema, nil
	}
	ema, err := iwc.indicator.EMA(pair, limit, period, from, to)
	if err != nil {
		return []float64{}, err
	}
	iwc.cache[pair] = ema
	return ema, nil
}
