package concurrency

import (
    "fmt"
    "testing"
)

func TestChannelFourthDemo(t *testing.T) {
    intChannel := make(chan int)

    go func(channel chan int) {
        for i := 0; i < 5; i++ {
            channel <- i
        }
        close(channel)
    }(intChannel)

    for i := 0; i < 6; i++ {
        value, open := <-intChannel
        if !open {
            break
        }
        fmt.Println(value)
    }
}
