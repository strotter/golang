// Multiple listeners on the same channel isn't ideal.
// Having a map and publishing to each works much better.

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

func create_market_timer(reader_map map[int]chan MarketDataBar) {
    ticker := time.NewTicker(time.Second)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at ", t)
            new_data := get_new_data()
            fmt.Println("New market data: ", new_data)

            for key, listening_chan := range reader_map {
                fmt.Println("Publishing to ", key)
                listening_chan <- new_data
            }
        }
    }()

    time.Sleep(time.Second * 10)
    ticker.Stop()
}

func create_listeners(reader_map map[int]chan MarketDataBar) {
    for i := 1; i <= 5; i++ {
        reader_map[i] = make(chan MarketDataBar)
    }

    for key, value := range reader_map {
        go func(number int, listening_chan <-chan MarketDataBar) {
            for v := range listening_chan {
                fmt.Println("Got data on channel %d: ", number, v)
            }
        }(key, value)
    }
}

func main() {
    reader_map := make(map[int]chan MarketDataBar)
    create_listeners(reader_map)
    create_market_timer(reader_map)

    fmt.Println("test")
}