package api

import (
	"fmt"
	"time"

	"github.com/hasan1786/gotraining/ticker"
	"github.com/parnurzeal/gorequest"
)

func GetBtcInfo(outdata chan map[string]ticker.Ticker, timeofRequest chan string) {

	fmt.Println("Get BTC Info")
	request := gorequest.New()

	var data map[string]ticker.Ticker

	_, _, errs := request.Get(BTC_INFO_URL).EndStruct(&data)

	//err := json.Unmarshal([]byte(body), &data)

	if errs != nil {

		fmt.Println("Error Parsing JSON", errs)
	}

	//timeOfRequest <- time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println("Response", timeOfRequest)
	//outdata <- 0
	outdata <- data
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Currency", "Buy", "Sell", "Symbol", "Current Time"})
	timeofRequest <- time.Now().Format("2006-01-02 15:04:05")
}
