package goroutine

import (
    "context"
    "fmt"
    "testing"
    "time"
)

func worker_controlled_by_context(ctx context.Context, id int, jobs <-chan int) {
    for {
        select {
        case job := <-jobs:
            fmt.Printf("Worker %d received job: %d\n", id, job)
            time.Sleep(500 * time.Millisecond)
        case <-ctx.Done():
            fmt.Printf("Worker %d exiting: %v\n", id, ctx.Err())
            return
        }
    }
}

func TestExitGoroutineByContext(t *testing.T) {
    jobChan := make(chan int)

    ctx, cancel := context.WithCancel(context.Background())

    go worker_controlled_by_context(ctx, 1, jobChan)
    go worker_controlled_by_context(ctx, 2, jobChan)

    for i := 1; i <= 5; i++ {
        jobChan <- i
    }

    time.Sleep(2 * time.Second)
    cancel()

    time.Sleep(1 * time.Second)
    fmt.Println("Main function exiting")
}
