package main

import (
	"encoding/json"
	"fmt"
)

type Bid struct {
	Price     string
	Size      string
	NumOrders int
}
type OrderBook struct {
	Sequence int64 `json:"sequence"`
	Bids     []Bid `json:"bids"`
	Asks     []Bid `json:"asks"`
}

func (b *Bid) UnmarshalJSON(data []byte) error {

	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	b.Price, _ = v[0].(string)
	b.Size, _ = v[1].(string)
	b.NumOrders = int(v[2].(float64))
	return nil
}

func main() {

	str := `{ 
		"sequence": 3977318850, 
		"bids": [ 
				  [ "4625.78", "0.80766325", 3 ] 
				], 
		"asks": [ 
				  [ "4625.79", "3.0154341", 3 ] 
				]
		}`

	var orders OrderBook
	err := json.Unmarshal([]byte(str), &orders)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orders)
}
