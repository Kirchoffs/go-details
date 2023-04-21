package reflect

import (
    "reflect"
    "testing"
)

func TestMapRelated(t *testing.T) {
    var m map[string]int
    typ := reflect.TypeOf(m)
    t.Log("Type of m: ", typ.String()) // Reflect.Type: map[string]int

    t.Log("Type of m's element (value): ", typ.Elem()) // Reflect.Type: int
    t.Log("Type of m's key: ", typ.Key())              // Reflect.Type: string
}
