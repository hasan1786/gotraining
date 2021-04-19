package main

import (
	"fmt"

	"github.com/hasan1786/gotraining/ticker"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_TICKER string = "https://blockchain.info/ticker"
)

func main() {

	request := gorequest.New()
	var data map[string]ticker.Ticker

	_, _, errs := request.Get(URL_TICKER).EndStruct(&data)

	//err := json.Unmarshal([]byte(body), &data)

	if errs != nil {

		fmt.Println("Error Parsing JSON", errs)
	}
	//fmt.Println("Response", resp)
	//fmt.Println("Response", body)

	for _, t := range data {
		t.PrintTickerInfo()
	}
}
