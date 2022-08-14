package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type binanceData [][]interface{}

type candle struct {
	symbol             string
	Open               float64
	Close              float64
	High               float64
	Low                float64
	interval           string
	Volum              float64
	UpperShadowpercent float64
	BottomSadowpercent float64
	OpenTime           int
	CloseTime          int
}

func convertToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func convertToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {

	for {
		getData("BTCUSDT", "1d")
		time.Sleep(time.Second * 5)
	}

}

func getData(symbol string, interval string) {
	response, err := http.Get("https://www.binance.com/api/v3/uiKlines?limit=1&symbol=" + symbol + "&interval=" + interval)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	data := binanceData{}

	d := json.NewDecoder(strings.NewReader(string(body)))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data[0])

	var candles []candle

	for _, v := range data {
		var candle candle
		candle.OpenTime = convertToInt(fmt.Sprint(v[0]))
		candle.Open = convertToFloat64(fmt.Sprint(v[1]))
		candle.Close = convertToFloat64(fmt.Sprint(v[4]))
		candle.High = convertToFloat64(fmt.Sprint(v[2]))
		candle.Low = convertToFloat64(fmt.Sprint(v[3]))
		candle.Volum = convertToFloat64(fmt.Sprint(v[5]))
		candle.CloseTime = convertToInt(fmt.Sprint(v[6]))
		candles = append(candles, candle)
		fmt.Println("price is:", candle.Close)
	}

}
