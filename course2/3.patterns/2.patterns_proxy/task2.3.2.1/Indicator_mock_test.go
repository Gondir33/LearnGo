package main

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	limit  = 30
	pair   = "BTC_USD"
	period = 5
)

func TestNewIndicator(t *testing.T) {
	exchange := NewMockExchanger(t)
	got := NewIndicator(exchange, WithEMA(calculateEMA), WithSMA(calculateSMA))
	want := &Indicator{exchange: exchange, calculateSMA: calculateSMA, calculateEMA: calculateEMA}
	assert.Equal(t, want.exchange, got.exchange)
}

func TestIndicator_SMA(t *testing.T) {
	type args struct {
		pair   string
		limit  int
		period int
		from   time.Time
		to     time.Time
	}
	exchange := NewMockExchanger(t)
	tests := []struct {
		name    string
		i       *Indicator
		args    args
		sma     []float64
		smaerr  error
		want    []float64
		wantErr bool
	}{{
		name:    "without error",
		i:       NewIndicator(exchange, WithEMA(calculateEMA), WithSMA(calculateSMA)),
		args:    args{pair: "BTC_USD", limit: 30, period: 5, from: time.Now().AddDate(0, 0, -1), to: time.Now()},
		sma:     []float64{100, 150, 250, 300, 350},
		smaerr:  nil,
		want:    []float64{230},
		wantErr: false,
	}, {
		name:    "with error",
		i:       NewIndicator(exchange, WithEMA(calculateEMA), WithSMA(calculateSMA)),
		args:    args{pair: "sadsada", limit: 30, period: 5, from: time.Now(), to: time.Now().AddDate(0, 0, -1)},
		sma:     []float64{},
		smaerr:  errors.New("Not response or smth"),
		want:    []float64{},
		wantErr: true,
	}}
	for _, tt := range tests {
		exchange.On("GetClosePrice", tt.args.pair, tt.args.limit, tt.args.from, tt.args.to).Return(tt.sma, tt.smaerr)
		got, err := tt.i.SMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
		if (err != nil) != tt.wantErr {
			t.Errorf("Indicator.SMA() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Indicator.SMA() = %v, want %v", got, tt.want)
		}
	}
}

func TestIndicator_EMA(t *testing.T) {
	type args struct {
		pair   string
		limit  int
		period int
		from   time.Time
		to     time.Time
	}
	exchange := NewMockExchanger(t)
	tests := []struct {
		name    string
		i       *Indicator
		args    args
		ema     []float64
		emaerr  error
		want    []float64
		wantErr bool
	}{{
		name:    "without error",
		i:       NewIndicator(exchange, WithEMA(calculateEMA), WithSMA(calculateSMA)),
		args:    args{pair: "BTC_USD", limit: 30, period: 5, from: time.Now().AddDate(0, 0, -1), to: time.Now()},
		ema:     []float64{100, 150, 250, 300, 350},
		emaerr:  nil,
		want:    []float64{100, 116.66666666666667, 161.11111111111111, 207.40740740740742, 254.93827160493828},
		wantErr: false,
	}, {
		name:    "with error",
		i:       NewIndicator(exchange, WithEMA(calculateEMA), WithSMA(calculateSMA)),
		args:    args{pair: "sadsada", limit: 30, period: 5, from: time.Now(), to: time.Now().AddDate(0, 0, -1)},
		ema:     []float64{},
		emaerr:  errors.New("Not response or smth"),
		want:    []float64{},
		wantErr: true,
	}}
	for _, tt := range tests {
		exchange.On("GetClosePrice", tt.args.pair, tt.args.limit, tt.args.from, tt.args.to).Return(tt.ema, tt.emaerr)
		got, err := tt.i.EMA(tt.args.pair, tt.args.limit, tt.args.period, tt.args.from, tt.args.to)
		if (err != nil) != tt.wantErr {
			t.Errorf("Indicator.EMA() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Indicator.EMA() = %v, want %v", got, tt.want)
		}
	}
}
