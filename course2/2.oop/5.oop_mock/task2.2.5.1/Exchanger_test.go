package main

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewExmo(t *testing.T) {
	client := &http.Client{}
	res := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(client))
	want := &Exmo{url: "https://api.exmo.com/v1.1", client: client}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("NewExmo res:%v, want:%v", res, want)
	}
}

func TestPrepareUrl(t *testing.T) {
	e := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{}))
	res := e.prepareUrl(ticker)
	want := "https://api.exmo.com/v1.1/ticker"
	if !reflect.DeepEqual(res, want) {
		t.Errorf("PrepareUrl res:%v, want:%v", res, want)
	}
}

func TestJoinPairs(t *testing.T) {
	e := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{}))
	got := e.joinPairs("1", "2")
	want := "1%2C2"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("JoinPairs res: %v, want: %v", got, want)
	}
}

func TestPrepareBody(t *testing.T) {
	e := NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{}))
	data := []struct {
		args []string
		want io.Reader
	}{{
		args: []string{},
		want: nil,
	}, {
		args: []string{"1", "2"},
		want: strings.NewReader("1&2"),
	}}
	for _, tt := range data {
		got := e.prepareBody(tt.args...)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("PrepareBody res: %v, want: %v", got, tt.want)
		}
	}
}

func TestExmo_GetTicker(t *testing.T) {
	tests := []struct {
		name    string
		exmo    *Exmo
		want    Ticker
		wantErr bool
	}{{
		name:    "1",
		exmo:    NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{})),
		want:    Ticker{},
		wantErr: false,
	}, {
		name:    "2",
		exmo:    NewExmo(WithURL("https://api.example.com/data"), WithClient(&http.Client{})),
		wantErr: true,
	}}
	for _, tt := range tests {
		e := &Exmo{
			client: tt.exmo.client,
			url:    tt.exmo.url,
		}
		got, err := e.GetTicker()
		if (err != nil) != tt.wantErr {
			t.Errorf("Exmo.GetTicker() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if reflect.DeepEqual(got, tt.want) {
			t.Errorf("Exmo.GetTicker() = %v, want %v", got, tt.want)
		}
	}
}

func TestExmo_GetTrades(t *testing.T) {
	type args struct {
		pairs []string
	}
	tests := []struct {
		name    string
		exmo    *Exmo
		args    args
		want    Trades
		wantErr bool
	}{{
		name:    "1",
		exmo:    NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{})),
		args:    args{[]string{"BTC_USD", "BTC_EUR"}},
		want:    Trades{},
		wantErr: false,
	}, {
		name:    "2",
		exmo:    NewExmo(WithURL("https://api.example.com/data"), WithClient(&http.Client{})),
		args:    args{[]string{"sadsadsd"}},
		wantErr: true,
	}}
	for _, tt := range tests {
		e := &Exmo{
			client: tt.exmo.client,
			url:    tt.exmo.url,
		}
		got, err := e.GetTrades(tt.args.pairs...)
		if (err != nil) != tt.wantErr {
			t.Errorf("Exmo.GetTrades() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if reflect.DeepEqual(got, tt.want) {
			t.Errorf("Exmo.GetTrades() = %v, want %v", got, tt.want)
		}
	}
}

func TestExmo_GetOrderBook(t *testing.T) {
	type args struct {
		limit int
		pairs []string
	}
	tests := []struct {
		name    string
		exmo    *Exmo
		args    args
		want    OrderBook
		wantErr bool
	}{{
		name:    "1",
		exmo:    NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{})),
		args:    args{limit: 100, pairs: []string{"BTC_USD", "BTC_EUR"}},
		want:    OrderBook{},
		wantErr: false,
	}, {
		name:    "3",
		exmo:    NewExmo(WithURL("https://api.example.com/data"), WithClient(&http.Client{})),
		args:    args{limit: 100, pairs: []string{"asdsadsadsa"}},
		wantErr: true,
	}}
	for _, tt := range tests {
		e := &Exmo{
			client: tt.exmo.client,
			url:    tt.exmo.url,
		}
		got, err := e.GetOrderBook(tt.args.limit, tt.args.pairs...)
		if (err != nil) != tt.wantErr {
			t.Errorf("Exmo.GetOrderBook() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if reflect.DeepEqual(got, tt.want) {
			t.Errorf("Exmo.GetOrderBook() = %v, want %v", got, tt.want)
		}
	}
}

func TestExmo_GetCurrencies(t *testing.T) {
	tests := []struct {
		name    string
		exmo    *Exmo
		want    Currencies
		wantErr bool
	}{{
		name:    "1",
		exmo:    NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{})),
		want:    Currencies{},
		wantErr: false,
	}, {
		name:    "2",
		exmo:    NewExmo(WithURL("https://api.example.com/data"), WithClient(&http.Client{})),
		wantErr: true,
	}}
	for _, tt := range tests {
		e := &Exmo{
			client: tt.exmo.client,
			url:    tt.exmo.url,
		}
		got, err := e.GetCurrencies()
		if (err != nil) != tt.wantErr {
			t.Errorf("Exmo.GetCurrencies() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if reflect.DeepEqual(got, tt.want) {
			t.Errorf("Exmo.GetCurrencies() = %v, want %v", got, tt.want)
		}
	}
}

func TestExmo_GetCandlesHistory(t *testing.T) {
	type args struct {
		pair  string
		limit int
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		exmo    *Exmo
		args    args
		want    CandlesHistory
		wantErr bool
	}{{
		name:    "1",
		exmo:    NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{})),
		args:    args{pair: "BTC_USD", limit: 30, start: time.Now().Add(-time.Hour * 24), end: time.Now()},
		want:    CandlesHistory{},
		wantErr: false,
	}, {
		name:    "2",
		exmo:    NewExmo(WithURL("\t"), WithClient(&http.Client{})),
		args:    args{},
		want:    CandlesHistory{[]Candle{{T: 1}}},
		wantErr: true,
	}}
	for _, tt := range tests {
		e := &Exmo{
			client: tt.exmo.client,
			url:    tt.exmo.url,
		}
		got, err := e.GetCandlesHistory(tt.args.pair, tt.args.limit, tt.args.start, tt.args.end)
		if (err != nil) != tt.wantErr {
			t.Errorf("Exmo.GetCandlesHistory() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if reflect.DeepEqual(got, tt.want) {
			t.Errorf("Exmo.GetCandlesHistory() = %v, want %v", got, tt.want)
		}
	}
}

func TestExmo_GetClosePrice(t *testing.T) {
	type args struct {
		pair  string
		limit int
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		exmo    *Exmo
		args    args
		want    []float64
		wantErr bool
	}{{
		name:    "1",
		exmo:    NewExmo(WithURL("https://api.exmo.com/v1.1"), WithClient(&http.Client{})),
		args:    args{pair: "BTC_USD", limit: 30, start: time.Now().Add(-time.Hour * 24), end: time.Now()},
		wantErr: false,
	}, {
		name:    "2",
		exmo:    NewExmo(WithURL("https://api.example.com/data"), WithClient(&http.Client{})),
		args:    args{},
		wantErr: true,
	}}
	for _, tt := range tests {
		e := &Exmo{
			client: tt.exmo.client,
			url:    tt.exmo.url,
		}
		got, err := e.GetClosePrice(tt.args.pair, tt.args.limit, tt.args.start, tt.args.end)
		if (err != nil) != tt.wantErr {
			t.Errorf("Exmo.GetClosePrice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		tt.want = got
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Exmo.GetClosePrice() = %v, want %v", got, tt.want)
		}
	}
}

func TestExmo_doRequest(t *testing.T) {
	e := NewExmo(WithURL(""), WithClient(&http.Client{}))
	_, err := e.doRequest("GET", "https://example.com", nil)
	if err != nil {
		t.Errorf("Exmo.doRequest wanterr but don't get")
	}
}
