package context

import (
    "context"
    "fmt"
    "testing"
    "time"
)

func TestNestedContext(t *testing.T) {
    ctx := context.Background()
    before := time.Now()
    parentCtx, _ := context.WithTimeout(ctx, 100*time.Millisecond)

    go func() {
        childCtx, _ := context.WithTimeout(parentCtx, 300*time.Millisecond)
        select {
        case <-childCtx.Done():
            after := time.Now()
            fmt.Println("child during:", after.Sub(before).Milliseconds())
        }
    }()

    select {
    case <-parentCtx.Done():
        after := time.Now()
        fmt.Println("parent during:", after.Sub(before).Milliseconds())
    }
}
