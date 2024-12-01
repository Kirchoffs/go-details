package context

import (
    "context"
    "fmt"
    "testing"
    "time"
)

func TestContextSecondDemo(t *testing.T) {
    ctx := context.Background()
    before := time.Now()
    preCtx, preCtxCancel := context.WithTimeout(ctx, 100*time.Millisecond)
    defer preCtxCancel()

    go func() {
        childCtx, childCtxCancel := context.WithTimeout(preCtx, 200*time.Millisecond)
        defer childCtxCancel()
        select {
        case <-childCtx.Done():
            after := time.Now()
            fmt.Println("child during:", after.Sub(before).Milliseconds())
        }
    }()

    select {
    case <-preCtx.Done():
        after := time.Now()
        fmt.Println("pre during:", after.Sub(before).Milliseconds())
    }
}
