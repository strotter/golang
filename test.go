package main

import "fmt"
import "math/rand"

type MarketData struct {
	price float64
}

type MarketDataBar struct {
	data MarketData
	high float64
	low float64
}

func (mkd *MarketDataBar) Randomize() {
	mkd.high = rand.Float64()
	mkd.low = rand.Float64()
}

func main() {
	z := 1.0
	fmt.Println(z)
	m := MarketData{price: 500.24}
	fmt.Println(m)
	// x := MarketDataBar{data: m, high: 500.99, low: 500.11}
	//fmt.Println(x)
	x := MarketDataBar{data: MarketData{price: 500.5}, high: 500.99, low: 500.11}
	fmt.Println(x)
	x.Randomize()
	fmt.Println(x)
}