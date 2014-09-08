package main

import (
    "fmt"
    "math/rand"
    "net"
    "bufio"
    "strconv"
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

func (mdb *MarketDataBar) String() string {
    return fmt.Sprintf("%f %f %f", mdb.data.price, mdb.high, mdb.low)
}

func get_new_data() MarketDataBar {
    mdb := MarketDataBar{data: MarketData{ price: rand.Float64() }, high: rand.Float64(), low: rand.Float64() }
    return mdb
}

const PORT = 3333

var listeners = make(map[net.Conn]chan string)

func main() {
    server, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))

    if server == nil {
        panic("Couldn't start listening: " + err.Error())
    }

    conns := client_conns(server)
    create_market_ticker()

    for {
        go handle_conn(<-conns)
    }
}

func client_conns(listener net.Listener) chan net.Conn {
    ch := make(chan net.Conn)
    i := 0

    go func() {
        for {
            client, err := listener.Accept()
            if client == nil {
                fmt.Println("Couldn't accept: " + err.Error())
                continue
            }

            i++
            fmt.Printf("%d: %v <-> %v", i, client.LocalAddr(), client.RemoteAddr())
            ch <- client
        }
    }()

    return ch
}

func create_market_ticker() {
    ticker := time.NewTicker(time.Second)

    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at ", t)
            new_data := get_new_data()
            fmt.Println("New market data: ", new_data)

            for key, listening_chan := range listeners {
                fmt.Println("Channel: ", key)
                string_data := new_data.String()
                fmt.Println("Publishing new data: ", string_data)
                listening_chan <- string_data
            }
        }
    }()
}

func add_to_market_ticker(client net.Conn, new_channel chan string) {
    fmt.Println("Adding to market ticker: ", new_channel)
    listeners[client] = new_channel
}

func handle_conn(client net.Conn) {
    fmt.Println("Handling connection: ", client)

    data_ready := make(chan string)

    go func() {
        b := bufio.NewReader(client)

        for {
            line, err := b.ReadString('\n')

            if err != nil {
                fmt.Println("Error on client: ", err.Error())
                delete(listeners, client)
                break
            }

            fmt.Println("Received: ", line)

            // Echo back to client
            data_ready <- line
        }
    }()

    add_to_market_ticker(client, data_ready)

    for {
        data := <-data_ready
        client.Write([]byte(data))
    }
}