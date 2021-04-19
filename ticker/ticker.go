package ticker

import "fmt"

type Ticker struct {
	O15m   float64 `json:"15m"`
	Last   float64 `json:"last"`
	Buy    float64 `json:"buy"`
	Sell   float64 `json:"sell"`
	Symbol string  `json:"symbol"`
}

func (t *Ticker) PrintTickerInfo() {

	if t != nil {
		fmt.Printf("Buy = %f%s , Sell = %f%s  \n", t.Buy, t.Symbol, t.Sell, t.Symbol)
	}
}
