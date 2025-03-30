package goroutine

import (
    "fmt"
    "testing"
    "time"
)

func worker_controlled_by_chan(id int, c <-chan int, done <-chan struct{}) {
    for {
        select {
        case job := <-c:
            fmt.Printf("Worker %d received job: %d\n", id, job)
            time.Sleep(500 * time.Millisecond)
        case <-done:
            fmt.Printf("Worker %d exiting\n", id)
            return
        }
    }
}

func TestExitGoroutineByChan(t *testing.T) {
    jobChan := make(chan int)
    done := make(chan struct{})

    go worker_controlled_by_chan(1, jobChan, done)
    go worker_controlled_by_chan(2, jobChan, done)

    for i := 1; i <= 5; i++ {
        jobChan <- i
    }

    time.Sleep(2 * time.Second)

    close(done)

    time.Sleep(1 * time.Second)
    fmt.Println("Main function exiting")
}
