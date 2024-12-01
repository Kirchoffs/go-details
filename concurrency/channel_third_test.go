package concurrency

import (
    "fmt"
    "testing"
)

func TestChannelThirdDemo(t *testing.T) {
    channel := make(chan string, 1)
    channel <- "Gift"
    fmt.Println(<-channel)
}
