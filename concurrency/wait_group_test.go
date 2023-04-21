package concurrency

import (
    "fmt"
    "sync"
    "testing"
)

func TestWaitGroupDemo(t *testing.T) {
    var beeper sync.WaitGroup
    ninjas := []string{"Ben", "Tom", "Jim"}
    beeper.Add(len(ninjas))

    for _, ninja := range ninjas {
        go attackWithWaitGroup(ninja, &beeper)
    }

    beeper.Wait()
    fmt.Println("Mission completed")
}

func attackWithWaitGroup(ninja string, beeper *sync.WaitGroup) {
    fmt.Println("Attact", ninja)
    beeper.Done()
}
