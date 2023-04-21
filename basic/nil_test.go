package basic

import (
    "fmt"
    "testing"
)

type MyInterface interface {
    GetValue() int
}

type MyStruct struct {
    Value int
}

func (x *MyStruct) GetValue() int {
    return x.Value
}

func TestNil(t *testing.T) {
    var ptrStruct *MyStruct = (*MyStruct)(nil)
    var ptrInterface MyInterface = ptrStruct    
    fmt.Println(ptrStruct)  
    fmt.Println(ptrInterface)   
    fmt.Println(ptrStruct == nil)    // true
    fmt.Println(ptrInterface != nil) // true
}
