package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dashboarder должен возвращать историю свечей и индикаторы с несколькими периодами, заданными через opts
type Dashboarder interface {
	GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error)
}

type DashboardData struct {
	Name           string
	CandlesHistory CandlesHistory
	Indicators     map[string][]IndicatorData
	limit          int
	from           time.Time
	to             time.Time
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Periods   []int
	Indicator Indicatorer
}

type Dashboard struct {
	exchange           Exchanger
	withCandlesHistory bool
	IndicatorOpts      []IndicatorOpt
	limit              int
	from               time.Time
	to                 time.Time
}

func (d *Dashboard) GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error) {
	for _, opt := range opts {
		opt(d)
	}

	candles, err := d.exchange.GetCandlesHistory(pair, d.limit, d.from, d.to)
	if err != nil {
		return DashboardData{}, err
	}

	indicators := make(map[string][]IndicatorData)
	for _, indicatorOpt := range d.IndicatorOpts {
		data := make([]IndicatorData, 0, len(indicatorOpt.Periods))
		var prices []float64
		for _, period := range indicatorOpt.Periods {
			if indicatorOpt.Name == "EMA" {
				prices, err = indicatorOpt.Indicator.EMA(pair, d.limit, period, d.from, d.to)
			} else if indicatorOpt.Name == "SMA" {
				prices, err = indicatorOpt.Indicator.SMA(pair, d.limit, period, d.from, d.to)
			}
			if err != nil {
				return DashboardData{}, err
			}
			data = append(data, IndicatorData{
				Name:     indicatorOpt.Name,
				Period:   period,
				Indicate: prices,
			})
		}
		indicators[indicatorOpt.Name] = data
	}
	return DashboardData{
		Name:           pair,
		CandlesHistory: candles,
		Indicators:     indicators,
		limit:          d.limit,
		from:           d.from,
		to:             d.to,
	}, nil
}

func WithCandlesHistory(limit int, from, to time.Time) func(*Dashboard) {
	return func(d *Dashboard) {
		d.limit = limit
		d.from = from
		d.to = to
	}
}

func WithIndicatorOpts(opts ...IndicatorOpt) func(*Dashboard) {
	return func(d *Dashboard) {
		d.IndicatorOpts = append(d.IndicatorOpts, opts...)
	}
}

func NewDashboard(exchange Exchanger) Dashboarder {
	return &Dashboard{exchange: exchange}
}

func main() {
	exchange := NewExmo(WithClient(&http.Client{}), WithURL("https://api.exmo.com/v1.1"))
	dashboard := NewDashboard(exchange)
	data, err := dashboard.GetDashboard("BTC_USD", WithCandlesHistory(30, time.Now().Add(-time.Hour*24*30), time.Now()), WithIndicatorOpts(
		IndicatorOpt{
			Name:      "SMA",
			Periods:   []int{5, 10, 20},
			Indicator: NewIndicator(exchange),
		},
		IndicatorOpt{
			Name:      "EMA",
			Periods:   []int{5, 10, 20},
			Indicator: NewIndicator(exchange),
		},
	))
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
