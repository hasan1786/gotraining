package utils

import "github.com/hasan1786/gotraining/ticker"

func CloseStringChan(queue chan string) {

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

func CloseMapChan(queue chan map[string]ticker.Ticker) {
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
