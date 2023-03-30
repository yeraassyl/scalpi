package main

import (
	"fmt"
	"os"
	"os/signal"
	"scalpi/market"
	"syscall"
)

func main() {
	//apiKey := ""

	doneC, stopC, err := market.WsDepthServe(url("bnbusdt"), depthDiffHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	select {
	case <-c:
		doneC <- struct{}{}
		stopC <- struct{}{}
	}
}

func depthDiffHandler(event *market.DepthEvent) {
	fmt.Println(event)
}

func errHandler(err error) {
	fmt.Println(err)
}

// TODO make wsDepthServe to receive symbol not url
func url(symbol string) string {
	return fmt.Sprintf("wss://stream.binance.com:9443/ws/%s@depth", symbol)
}
