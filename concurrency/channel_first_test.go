package concurrency

import (
    "fmt"
    "math/rand"
    "testing"
    "time"
)

func TestChannelFirstDemo(t *testing.T) {
    channel := make(chan string)
    numRounds := 3
    go throwingNinjaStar(channel, numRounds)
    for i := 0; i < numRounds; i++ {
        fmt.Println(<-channel)
    }
}

func throwingNinjaStar(channel chan string, numRounds int) {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < numRounds; i++ {
        score := rand.Intn(10)
        channel <- fmt.Sprint("Your scored: ", score)
    }
}
