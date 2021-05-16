package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

type MakerTaker struct {
	Maker float64 `json:"maker"`
	Taker float64 `json:"taker"`
}

type Fee struct {
	Bid MakerTaker `json:"bid"`
	Ask MakerTaker `json:"ask"`
}

func (f *Fee) UnmarshalJSON(data []byte) error {
	var aux map[string]interface{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	mapstructure.Decode(aux, f)
	return nil
}

type Market struct {
	BaseMarket         string    `json:"baseMarket"`
	QuoteMarket        string    `json:"quoteMarket"`
	MinBuyAmount       float64   `json:"minBuyAmount,omitempty"`
	MinSellAmount      float64   `json:"minSellAmount,omitempty"`
	BasePrecision      int       `json:"basePrecision,omitempty"`
	QuotePrecision     int       `json:"quotePrecision,omitempty"`
	Status             string    `json:"status"`
	Fee                Fee       `json:"fee,omitempty"`
	Low                float64   `json:"low,omitempty"`
	High               float64   `json:"high,omitempty"`
	Last               float64   `json:"last,omitempty"`
	Type               string    `json:"type"`
	Open               float64   `json:"open,omitempty"`
	Volume             float64   `json:"volume,omitempty"`
	Sell               float64   `json:"sell,omitempty"`
	Buy                float64   `json:"buy,omitempty"`
	At                 time.Time `json:"at,omitempty"`
	MaxBuyAmount       float64   `json:"maxBuyAmount,omitempty"`
	MinBuyVolume       float64   `json:"minBuyVolume,omitempty"`
	MaxBuyVolume       float64   `json:"maxBuyVolume,omitempty"`
	FeePercentOnProfit float64   `json:"feePercentOnProfit,omitempty"`
}

func (m *Market) UnmarshalJSON(data []byte) error {
	aux := &struct {
		BaseMarket         string  `json:"baseMarket"`
		QuoteMarket        string  `json:"quoteMarket"`
		MinBuyAmount       float64 `json:"minBuyAmount,omitempty"`
		MinSellAmount      float64 `json:"minSellAmount,omitempty"`
		BasePrecision      int     `json:"basePrecision,omitempty"`
		QuotePrecision     int     `json:"quotePrecision,omitempty"`
		Status             string  `json:"status"`
		At                 int64   `json:"at,omitempty"`
		Low                string  `json:"low,omitempty"`
		High               string  `json:"high,omitempty"`
		Last               string  `json:"last,omitempty"`
		Volume             string  `json:"volume,omitempty"`
		Sell               string  `json:"sell,omitempty"`
		Buy                string  `json:"buy,omitempty"`
		MaxBuyAmount       float64 `json:"maxBuyAmount,omitempty"`
		MinBuyVolume       float64 `json:"minBuyVolume,omitempty"`
		MaxBuyVolume       float64 `json:"maxBuyVolume,omitempty"`
		FeePercentOnProfit float64 `json:"feePercentOnProfit,omitempty"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	m.Low, _ = strconv.ParseFloat(aux.Low, 64)
	m.High, _ = strconv.ParseFloat(aux.High, 64)
	m.Last, _ = strconv.ParseFloat(aux.Last, 64)
	m.Volume, _ = strconv.ParseFloat(aux.Volume, 64)
	m.Sell, _ = strconv.ParseFloat(aux.Sell, 64)
	m.Buy, _ = strconv.ParseFloat(aux.Buy, 64)
	m.At = time.Unix(aux.At, 0)
	return nil
}

type Asset struct {
	Type              string  `json:"type"`
	Name              string  `json:"name"`
	Deposit           string  `json:"deposit"`
	Withdrawal        string  `json:"withdrawal"`
	ListingType       string  `json:"listingType"`
	Category          string  `json:"category"`
	WithdrawFee       float64 `json:"withdrawFee,omitempty"`
	MinWithdrawAmount float64 `json:"minWithdrawAmount,omitempty"`
	MaxWithdrawAmount float64 `json:"maxWithdrawAmount,omitempty"`
	MinDepositAmount  float64 `json:"minDepositAmount,omitempty"`
	Confirmations     int     `json:"confirmations,omitempty"`
}

type MarketStatus struct {
	Markets []Market `json:"markets"`
	Assets  []Asset  `json:"assets"`
}

// MarketStatus returs overview of markets and assets
func marketStatus() (data MarketStatus, err error) {
	resp, err := http.Get("https://api.wazirx.com/api/v2/market-status")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MarketStatus{}, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return MarketStatus{}, err
	}
	return data, nil
}

func main() {
	data, err := marketStatus()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
