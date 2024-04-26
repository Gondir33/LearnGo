package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/guptarohit/asciigraph"
)

type Pairs struct {
	BtcUsd Usd `json:"BTC_USD"`
	LtcUsd Usd `json:"LTC_USD"`
	EthUsd Usd `json:"ETH_USD"`
}

type Usd struct {
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

var stopDrow = make(chan struct{}, 1)
var Data []float64 = make([]float64, 0, 100)
var flag bool = false

func clearTerminal() {
	fmt.Println("\033[H\033[2J")
}

func mainMenu() {
	clearTerminal()
	fmt.Println("1. BTC_USD")
	fmt.Println("2. LTC_USD")
	fmt.Println("3. ETH_USD")
	fmt.Println("\nPress 1-3 to change symbol, press q to exit")
}

func handleButtons(char rune, key keyboard.Key) {
	if char == '1' && flag == false {
		startDrowGraph(char)
	} else if char == '2' && flag == false {
		startDrowGraph(char)
	} else if char == '3' && flag == false {
		startDrowGraph(char)
	} else if char == '\x00' && key == 127 && flag == true {
		flag = false
		stopDrow <- struct{}{}
		mainMenu()
	} else if char == 'q' {
		stopDrow <- struct{}{}
		os.Exit(0)
	}
}

// точка входа
func main() {

	mainMenu()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Printf("keyboard %v", err)
			os.Exit(1)
		}
		handleButtons(char, key)
	}
}

func startDrowGraph(char rune) {
	flag = true
	clearTerminal()
	Data = Data[:0]
	go func() {
		for {
			select {
			case <-stopDrow:
				return
			default:
				updateGraph(char)
				time.Sleep(time.Second)
			}
		}
	}()
}

func updateGraph(char rune) {
	data := getPrice()
	var caption string
	if char == '1' {
		s, err := strconv.ParseFloat(data.BtcUsd.LastTrade, 64)
		if err != nil {
			fmt.Printf("parse float %v", err)
			os.Exit(1)
		}
		Data = append(Data, s)
		caption = "BTC_USD " + data.BtcUsd.LastTrade
	} else if char == '2' {
		s, err := strconv.ParseFloat(data.LtcUsd.LastTrade, 64)
		if err != nil {
			fmt.Printf("parse float %v", err)
			os.Exit(1)
		}
		Data = append(Data, s)
		caption = "LTC_USD " + data.LtcUsd.LastTrade
	} else {
		s, err := strconv.ParseFloat(data.EthUsd.LastTrade, 64)
		if err != nil {
			fmt.Printf("parse float %v", err)
			os.Exit(1)
		}
		Data = append(Data, s)
		caption = "ETH_USD " + data.EthUsd.LastTrade
	}
	graph := asciigraph.Plot(Data, asciigraph.Width(100),
		asciigraph.Height(10),
		asciigraph.CaptionColor(asciigraph.Green),
		asciigraph.Caption(caption), asciigraph.LabelColor(asciigraph.Red),
	)
	clearTerminal()
	fmt.Println(graph)
	currentTime := time.Now()
	currentTimeFormatted := currentTime.Format("15:04:05")
	currentDateFormatted := currentTime.Format("2006-01-02")
	fmt.Printf("Текущая дата: %s\n", currentDateFormatted)
	fmt.Printf("Текущее время: %s\n", currentTimeFormatted)
}

func getPrice() Pairs {
	var req *http.Request
	var err error

	client := &http.Client{}
	req, err = http.NewRequest("GET", "https://api.exmo.com/v1.1/ticker", nil)

	if err != nil {
		clearTerminal()
		fmt.Printf("error request %v", err)
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error response %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error read response %v", err)
		os.Exit(1)
	}
	var res Pairs
	json.Unmarshal(body, &res)
	return res
}
