package ticker

import "fmt"

type Ticker struct {
	O15m   float64 `json:"15m"`
	Last   float64 `json:"last"`
	Buy    float64 `json:"buy"`
	Sell   float64 `json:"sell"`
	Symbol string  `json:"symbol"`
}

func (t *Ticker) PrintTickerInfo(key string) {

	if t != nil {
		fmt.Printf("Currency = %s Buy = %f%s , Sell = %f%s  \n", key, t.Buy, t.Symbol, t.Sell, t.Symbol)
	}
}
