package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type OrderBookPair struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}
type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

type Pair struct {
	TradeID  int64  `json:"trade_id"`
	Date     int64  `json:"date"`
	Type     string `json:"type"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
}

type Currencies []string

type OrderBook map[string]OrderBookPair

type Ticker map[string]TickerValue

type Trades map[string][]Pair

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	exmo := &Exmo{}
	for _, opt := range opts {
		opt(exmo)
	}
	return exmo
}

func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}

func WithURL(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

func (e *Exmo) prepareUrl(api string) string {
	return e.url + api
}

func (e *Exmo) joinPairs(pairs ...string) string {
	return strings.Join(pairs, "%2C")
}

func (e *Exmo) prepareBody(params ...string) io.Reader {
	if len(params) == 0 {
		return nil
	}
	return strings.NewReader(strings.Join(params, "&"))
}

func (e *Exmo) doRequest(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bodyText, nil
}

func (e *Exmo) execute(api, method string, structPointer any, params ...string) error {
	url := e.prepareUrl(api)
	body := e.prepareBody(params...)
	bodyText, err := e.doRequest(method, url, body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyText, structPointer)
}

func (e *Exmo) GetTicker() (Ticker, error) {
	res := make(Ticker)
	err := e.execute(ticker, "POST", &res)
	if err != nil {
		return Ticker{}, err
	}
	return res, nil
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	res := make(Trades)
	err := e.execute(trades, "POST", &res, fmt.Sprintf("pair=%v", e.joinPairs(pairs...)))
	if err != nil {
		return Trades{}, err
	}
	return res, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	res := make(OrderBook)
	appendApi := make([]string, 2, 2)
	appendApi[0] = fmt.Sprintf("pair=%v", e.joinPairs(pairs...))
	appendApi[1] = fmt.Sprintf("limit=%v", limit)
	err := e.execute(orderBook, "POST", &res, appendApi...)
	if err != nil {
		return OrderBook{}, err
	}
	return res, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	var res Currencies
	err := e.execute(currency, "POST", &res)
	if err != nil {
		return Currencies{}, err
	}
	return res, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	var res CandlesHistory
	appendApi := make([]string, 4, 4)
	appendApi[0] = fmt.Sprintf("?symbol=%v", pair)
	appendApi[1] = fmt.Sprintf("resolution=%v", limit)
	appendApi[2] = fmt.Sprintf("from=%v", start.Unix())
	appendApi[3] = fmt.Sprintf("to=%v", end.Unix())
	err := e.execute(candlesHistory+strings.Join(appendApi, "&"), "GET", &res)
	if err != nil {
		return CandlesHistory{}, err
	}
	return res, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	candlehistory, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		return []float64{}, err
	}
	res := make([]float64, 0, len(candlehistory.Candles))
	for _, candle := range candlehistory.Candles {
		res = append(res, candle.C)
	}
	return res, nil
}

/*
func main() {
	var exchange Exchanger
	exchange = NewExmo(WithClient(&http.Client{}), WithURL("https://api.exmo.com/v1.1"))
	ticker, err := exchange.GetClosePrice("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ticker)
}
*/
