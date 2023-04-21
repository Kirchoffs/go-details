package reflect

import (
    "reflect"
    "testing"
)

func TestBasicPointer(t *testing.T) {
    x := 42
    px := &x

    rv := reflect.ValueOf(px)
    t.Logf("Kind: %s", rv.Kind()) // ptr

    rvIndirect := reflect.Indirect(rv)
    t.Logf("Kind: %s", rvIndirect.Kind()) // int
    t.Logf("Value: %d", rvIndirect.Int()) // 42
}
