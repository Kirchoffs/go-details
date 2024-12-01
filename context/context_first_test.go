package context

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "testing"
    "time"
)

func TestContextFirstDemo(t *testing.T) {
    timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
    defer cancel()

    req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://via.placeholder.com/2000x2000", nil)
    if err != nil {
        panic(err)
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    imageData, err := io.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    fmt.Printf("download image of size %d\n", len(imageData))
}
