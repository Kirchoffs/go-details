package reflect

import (
    "fmt"
    "reflect"
    "testing"
)

type Duck struct {
    name string
}

func (d Duck) Quack() string {
    return "Quack!"
}

func (pd *Duck) QuackWisely() int {
    fmt.Println("Quack wisely!")
    return 42
}

func (pd *Duck) QuackWith(content string) string {
    fmt.Println("Quack: " + content)
    return "Quack: " + content
}

func TestStructMethod(t *testing.T) {
    d := Duck{name: "Donald"}
    dType := reflect.TypeOf(d)
    dMethods := dType.NumMethod()
    for i := 0; i < dMethods; i++ {
        method := dType.Method(i)
        methodType := method.Type
        methodName := method.Name
        t.Logf("Method Name: %s, Type: %s", methodName, methodType)
    }

    pd := &Duck{name: "Daisy"}
    pdType := reflect.TypeOf(pd)
    pdMethods := pdType.NumMethod()
    for i := 0; i < pdMethods; i++ {
        method := pdType.Method(i)

        methodPkgPath := method.PkgPath
        methodName := method.Name
        methodType := method.Type

        t.Logf("Method Name: %s, PkgPath: %s", methodName, methodPkgPath)
        for j := 0; j < methodType.NumIn(); j++ {
            t.Logf("Method %s, In %d: %s", methodName, j, methodType.In(j))
        }
        for j := 0; j < methodType.NumOut(); j++ {
            t.Logf("Method %s, Out %d: %s", methodName, j, methodType.Out(j))
        }
    }
}
