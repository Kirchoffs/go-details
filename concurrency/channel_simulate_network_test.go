package concurrency

import (
    "fmt"
    "testing"
    "time"
)

type Network struct {
    ch   chan string
    done chan bool
}

func sendRequest(network *Network, req string) bool {
    select {
    case network.ch <- req:
        fmt.Println("Request sent:", req)
        return true
    case <-network.done:
        fmt.Println("Network is destroyed. Request not sent.")
        return false
    }
}

func destroyNetwork(network *Network) {
    fmt.Println("Destroying network...")
    close(network.done)
}

func TestChannelSimulateNetwork(t *testing.T) {
    network := &Network{
        ch:   make(chan string, 1),
        done: make(chan bool),
    }

    go func() {
        for req := range network.ch {
            fmt.Println("Processing request:", req)
            time.Sleep(1 * time.Second)
        }
    }()

    req := "GET /api/data"
    if sendRequest(network, req) {
        fmt.Println("Request was successfully sent.")
    } else {
        fmt.Println("Failed to send request.")
    }

    destroyNetwork(network)

    req = "GET /api/status"
    if sendRequest(network, req) {
        fmt.Println("Request was successfully sent.")
    } else {
        fmt.Println("Failed to send request.")
    }

    time.Sleep(2 * time.Second)
    fmt.Println("Program exited.")
}
