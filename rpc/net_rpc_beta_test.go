package rpc

import (
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

func TestRPCBeta(t *testing.T) {
    go startServerBeta(t)

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

func startServerBeta(t *testing.T) {
    arith := new(Arith)

    rpc.Register(arith)

    listener, err := net.Listen("tcp", ":6174")
    if err != nil {
        t.Error("Listen error:", err)
        return
    }
    defer listener.Close()

    t.Log("Server is running on port 6174...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            t.Log("Accept error:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
