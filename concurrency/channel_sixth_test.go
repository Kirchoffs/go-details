package concurrency

import (
    "fmt"
    "testing"
)

func TestChannelSixthDemo(t *testing.T) {
    ch := createChannel()

    for i := 0; i < 5; i++ {
        fmt.Println(<-ch)
    }
}

func createChannel() <-chan int {
    ch := make(chan int, 5)

    for i := 0; i < 5; i++ {
        ch <- i
    }

    return ch
}
