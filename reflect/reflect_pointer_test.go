package reflect

import (
    "reflect"
    "testing"
)

func TestBasicPointer(t *testing.T) {
    x := 42
    px := &x

    reflectValue := reflect.ValueOf(px)
    t.Logf("Kind: %s", reflectValue.Kind()) // ptr

    reflectIndirect := reflect.Indirect(reflectValue)
    t.Logf("Kind: %s", reflectIndirect.Kind()) // int
    t.Logf("Value: %d", reflectIndirect.Int()) // 42
}
