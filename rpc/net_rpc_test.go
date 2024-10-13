package rpc

import (
    "fmt"
    "net"
    "net/rpc"
    "testing"
    "time"
)

type Arith int

type Args struct {
    A, B int
}

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func TestRPC(t *testing.T) {
    go startServer()

    time.Sleep(time.Second)

    client, err := rpc.Dial("tcp", "localhost:6174")
    if err != nil {
        t.Fatalf("Dialing error: %v", err)
    }

    args := Args{A: 7, B: 8}
    var reply int

    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        t.Fatalf("RPC call error: %v", err)
    }

    expected := 56
    if reply != expected {
        t.Errorf("Expected %d, got %d", expected, reply)
    }
}

func startServer() {
    arith := new(Arith)

    rpc.Register(arith)

    l, err := net.Listen("tcp", ":6174")
    if err != nil {
        fmt.Println("Listen error:", err)
        return
    }
    defer l.Close()

    fmt.Println("Server is running on port 6174...")

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Accept error:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
