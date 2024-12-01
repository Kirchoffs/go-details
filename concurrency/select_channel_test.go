package concurrency

import (
    "fmt"
    "testing"
)

func TestSelectChannelDemo(t *testing.T) {
    ninjaAlpha, ninjaBeta := make(chan string), make(chan string)

    go captainElect(ninjaAlpha, "Ninja alpha")
    go captainElect(ninjaBeta, "Ninja beta")

    select {
    case message := <-ninjaAlpha:
        fmt.Println(message)
    case message := <-ninjaBeta:
        fmt.Println(message)
    }
}

func captainElect(channel chan string, message string) {
    channel <- message
}
