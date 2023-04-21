package rpc

import (
    "fmt"
    "net"
    "net/rpc"
    "testing"
    "time"
)

type HelloService struct{}

func (service *HelloService) Hello(request string, reply *string) error {
    *reply = "hello:" + request
    return nil
}

func TestRPCAlpha(t *testing.T) {
    go startServerAlpha(t)
    time.Sleep(time.Second)

    client, err := rpc.Dial("tcp", "localhost:6174")
    if err != nil {
        t.Fatalf("Dialing error: %v", err)
    }

    var reply string
    err = client.Call("HelloServiceAlpha.Hello", "hello", &reply)
    if err != nil {
        t.Fatalf("RPC call error: %v", err)
    }

    expected := "hello:hello"
    if reply != expected {
        t.Errorf("Expected %s, got %s", expected, reply)
    }
}

func startServerAlpha(t *testing.T) {
    rpc.RegisterName("HelloServiceAlpha", new(HelloService)) // Register the service

    listener, err := net.Listen("tcp", ":6174")
    if err != nil {
        fmt.Println("Listen error:", err)
        return
    }
    defer listener.Close()

    conn, err := listener.Accept()
    if err != nil {
        t.Error("Accept error:", err)
    }

    rpc.ServeConn(conn) // Serve the connection with the registered service
}
