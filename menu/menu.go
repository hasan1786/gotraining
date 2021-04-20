package menu

import (
	"fmt"
	"sync"
	"time"

	"github.com/hasan1786/gotraining/api"
	"github.com/hasan1786/gotraining/ticker"
	"github.com/hasan1786/gotraining/utils"
	"github.com/hasan1786/gotraining/writers"
)

var (
	quit          = make(chan int, 1)
	output        = make(chan map[string]ticker.Ticker, 1)
	timeofRequest = make(chan string, 1)
	timeouts      = time.NewTicker(10 * time.Second)
)

func quiting(quit chan int) {
	fmt.Scanln()
	quit <- 1
	fmt.Println("Quiting")
}
func start(wg *sync.WaitGroup) {

	for {

		select {
		case x := <-output:
			fmt.Println("Table writing")
			go writers.WriteToTable(x, <-timeofRequest)
		case <-timeouts.C:
			go api.GetBtcInfo(output, timeofRequest)
		case <-quit:
			timeouts.Stop()
			close(quit)
			fmt.Println("quit close")
			utils.CloseMapChan(output)
			fmt.Println("map close")
			utils.CloseStringChan(timeofRequest)
			fmt.Println("timeout close")
			wg.Done()
			return
		}
		time.Sleep(time.Millisecond * 100)
	}
}
func Init() {

	fmt.Println("Press Enter to close the program any time")

	var wg sync.WaitGroup
	wg.Add(1)

	go start(&wg)
	go quiting(quit)
	wg.Wait()
}
