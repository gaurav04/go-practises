package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TickerData struct {
	BaseUnit  string    `json:"base_unit"`
	QuoteUnit string    `json:"quote_unit"`
	Low       float64   `json:"low"`
	High      float64   `json:"high"`
	Last      float64   `json:"last"`
	Type      string    `json:"type"`
	Open      float64   `json:"open"`
	Volume    float64   `json:"volume"`
	Sell      float64   `json:"sell"`
	Buy       float64   `json:"buy"`
	At        time.Time `json:"at"`
	Name      string    `json:"name"`
}

func (t *TickerData) UnmarshalJSON(data []byte) error {

	aux := struct {
		At     int64  `json:"at"`
		Low    string `json:"low,omitempty"`
		High   string `json:"high,omitempty"`
		Last   string `json:"last,omitempty"`
		Volume string `json:"volume,omitempty"`
		Sell   string `json:"sell,omitempty"`
		Buy    string `json:"buy,omitempty"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Low, _ = strconv.ParseFloat(aux.Low, 64)
	t.High, _ = strconv.ParseFloat(aux.High, 64)
	t.Last, _ = strconv.ParseFloat(aux.Last, 64)
	t.Volume, _ = strconv.ParseFloat(aux.Volume, 64)
	t.Sell, _ = strconv.ParseFloat(aux.Sell, 64)
	t.Buy, _ = strconv.ParseFloat(aux.Buy, 64)
	t.At = time.Unix(aux.At, 0)
	return nil
}

func MarketTicker() (data map[string]TickerData, err error) {
	resp, err := http.Get("https://api.wazirx.com/api/v2/tickers")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	data, err := MarketTicker()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range data {
		fmt.Println(k, v)
	}
}
