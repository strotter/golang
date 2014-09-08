// Attempt at multiple goroutines listening on the same channel.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MarketData struct {
	price float64
}

type MarketDataBar struct {
	data MarketData
	high float64
	low float64
}

func get_new_data() MarketDataBar {
	mdb := MarketDataBar{data: MarketData{ price: rand.Float64() }, high: rand.Float64(), low: rand.Float64() }
	return mdb
}

func create_market_timer(writer chan<- MarketDataBar) {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
			new_data := get_new_data()
			fmt.Println("New market data: ", new_data)
			writer <- new_data
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop()
}

func create_listeners(reader <-chan MarketDataBar) {
	for i := 1; i <= 5; i++ {
		go func(j int, r <-chan MarketDataBar) {
			for v := range r {
				fmt.Println("Got data on %d: ", j, v)
			}
		}(i, reader)
	}
}

func main() {
	market_chan := make(chan MarketDataBar)
	create_market_timer(market_chan)
	create_listeners(market_chan)
	close(market_chan)

	fmt.Println("test")
}