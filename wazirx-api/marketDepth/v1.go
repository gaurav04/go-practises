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

type PriceVolume struct {
	PRICE  float64
	VOLUME float64
}

type MarketDepth struct {
	Timestamp time.Time     `json:"timestamp"`
	Asks      []PriceVolume `json:"asks"`
	Bids      []PriceVolume `json:"bids"`
}

func (d *MarketDepth) UnmarshalJSON(data []byte) error {
	aux := struct {
		Timestamp int64      `json:"timestamp"`
		Asks      [][]string `json:"asks"`
		Bids      [][]string `json:"bids"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var tmp PriceVolume
	for _, ask := range aux.Asks {
		tmp.PRICE, _ = strconv.ParseFloat(ask[0], 64)
		tmp.VOLUME, _ = strconv.ParseFloat(ask[1], 64)
		d.Asks = append(d.Asks, tmp)
	}

	for _, bid := range aux.Bids {
		tmp.PRICE, _ = strconv.ParseFloat(bid[0], 64)
		tmp.VOLUME, _ = strconv.ParseFloat(bid[1], 64)
		d.Bids = append(d.Bids, tmp)
	}

	d.Timestamp = time.Unix(aux.Timestamp, 0)
	return nil
}

func MarketDepths(market string) (data MarketDepth, err error) {
	url := fmt.Sprintf("https://api.wazirx.com/api/v2/depth?market=%s", market)
	fmt.Println(url)
	resp, err := http.Get(url)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MarketDepth{}, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return MarketDepth{}, err
	}
	return data, nil
}

func main() {
	s := "batinr"
	data, err := MarketDepths(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(data)

}
