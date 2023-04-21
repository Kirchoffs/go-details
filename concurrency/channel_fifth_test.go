package concurrency

import (
    "fmt"
    "testing"
)

func TestChannelFifthDemo(t *testing.T) {
    intChannel := make(chan int)

    go func(channel chan int) {
        for i := 0; i < 5; i++ {
            channel <- i
        }
        close(channel)
    }(intChannel)

    for value := range intChannel {
        fmt.Println(value)
    }
}
