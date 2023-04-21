package concurrency

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

var (
    rwLock sync.RWMutex
)

func TestRWMutextFirstDemo(t *testing.T) {
    readAndWrite()
}

func readAndWrite() {
    go read()
    go read()
    go read()
    go write()

    time.Sleep(5 * time.Second)
    fmt.Println("Done")
}

func read() {
    rwLock.RLock()
    defer rwLock.RUnlock()

    fmt.Println("Read locking")
    time.Sleep(time.Second)
    fmt.Println("Read unlocking")
}

func write() {
    rwLock.Lock()
    defer rwLock.Unlock()

    fmt.Println("Write locking")
    time.Sleep(time.Second)
    fmt.Println("Write unlocking")
}
