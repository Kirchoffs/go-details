package codec

import (
    "bytes"
    "fmt"
    "testing"
)

func TestBytesBasic(t *testing.T) {
    var buf bytes.Buffer

    buf.WriteString("Hello, ")
    buf.Write([]byte("world!"))
    fmt.Println("Buffer after writes:", buf.String())

    buf.WriteByte(' ')
    buf.Write([]byte("This is a test!"))
    fmt.Println("Buffer after more writes:", buf.String())

    data := buf.Bytes()
    fmt.Println("Buffer as byte slice:", data)

    buf.Reset()
    fmt.Println("Length of buffer after reset:", len(buf.Bytes()))
}
