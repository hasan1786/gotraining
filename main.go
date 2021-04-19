package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/hasan1786/gotraining/ticker"
	"github.com/olekukonko/tablewriter"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_TICKER string = "https://blockchain.info/ticker"
)

func callGetRequest(outdata chan map[string]ticker.Ticker, timeofRequest chan string) {

	request := gorequest.New()

	var data map[string]ticker.Ticker

	_, _, errs := request.Get(URL_TICKER).EndStruct(&data)

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
func main() {

	fmt.Println("Press Enter to close the program any time")

	var wg sync.WaitGroup
	quit := make(chan int, 1)
	output := make(chan map[string]ticker.Ticker, 1)
	timeofRequest := make(chan string, 1)
	//for currency, t := range data {
	//b, err := json.Marshal(t)
	//	table.([] string {currency, fmt.Sprintf("%f", t.Buy), fmt.Sprintf("%f", t.Sell), t.Symbol, timestamp} )
	//append(data, append([]string(b), curret)
	//}
	//fmt.Println("Response", resp)
	//fmt.Println("Response", body)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	timeouts := time.NewTicker(10 * time.Second)
	wg.Add(1)
	go func() {

		for {
			select {
			case x := <-output:
				fmt.Println("Table writing")
				go WriteToTable(table, x, <-timeofRequest)
			case <-timeouts.C:
				callGetRequest(output, timeofRequest)
			case <-quit:
				timeouts.Stop()

				close(quit)
				fmt.Println("quit close")
				closeBufferedChannelMap(output)
				fmt.Println("map close")
				closeBufferedChannelString(timeofRequest)
				fmt.Println("timeout close")
				wg.Done()
				return
				//close(quit)
			}
			time.Sleep(time.Millisecond * 100)
		}

	}()
	go func() {
		fmt.Scanln()
		quit <- 1
		fmt.Println("Quiting")
	}()
	wg.Wait()

	//fmt.Scanln()

}

func WriteToTable(table *tablewriter.Table, data map[string]ticker.Ticker, timestamp string) {

	table.ClearRows()
	for currency, t := range data {
		//b, err := json.Marshal(t)
		table.Append([]string{currency, fmt.Sprintf("%f", t.Buy), fmt.Sprintf("%f", t.Sell), t.Symbol, timestamp})
		//append(data, append([]string(b), curret)
	}

	table.Render()

}

func closeBufferedChannelString(queue chan string) {

	defer close(queue)
	for {
		select {
		case <-queue:
			continue
		default:
			//ok = false
			return
		}

	}
}

func closeBufferedChannelMap(queue chan map[string]ticker.Ticker) {
	defer close(queue)
	for {
		select {
		case <-queue:
			continue
		default:
			return
		}

	}

}
