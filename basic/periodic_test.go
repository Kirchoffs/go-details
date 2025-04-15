package basic

import (
    "testing"
    "time"
)

func TestTimeSleep(t *testing.T) {
    tickerInterval := 1 * time.Second

    cnt := 5
    for cnt > 0 {
        t.Logf("Task executed at: %s", time.Now())
        time.Sleep(tickerInterval)
        cnt--
    }
}

func TestTimeTicker(t *testing.T) {
    tickerInterval := 1 * time.Second

    timer := time.NewTimer(tickerInterval)
    defer timer.Stop()

    cnt := 5
    for cnt > 0 {
        select {
        case <-timer.C:
            t.Logf("Task executed at: %s", time.Now())
            timer.Reset(tickerInterval)
            cnt--
        }
    }
}
