package main

import (
	"encoding/json"
	"fmt"

	"github.com/hasan1786/gotraining/ticker"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_TICKER string = "https://blockchain.info/ticker"
)

func main() {

	request := gorequest.New()

	_, body, _ := request.Get(URL_TICKER).End()

	var data map[string]ticker.Ticker
	err := json.Unmarshal([]byte(body), &data)

	if err != nil {

		fmt.Println("Error Parsing JSON", err)
	}
	for _, t := range data {
		t.PrintTickerInfo()
	}
}
