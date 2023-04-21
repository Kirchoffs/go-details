package reflect

import (
    "reflect"
    "testing"
)

// In Go, basic types (such as `int`, `string`, `bool`, etc.) are predefined named types.
// Although they are not explicitly defined using the `type` keyword, they are assigned names in the language specification.
// Therefore, `reflect.TypeOf(x).Name()` will return their names.

// Unnamed types typically refer to the following cases:
// - Anonymous structs: For example, `struct { Name string }`.
// - Composite types: For example, `[]int` (slice), `map[string]int` (map), `chan int` (channel), etc.
// - Function types: For example, `func(int) int`.
// - Interface types: For example, `interface{}`.
// For these unnamed types, `reflect.TypeOf(x).Name()` will return an empty string.

func TestNameOfType(t *testing.T) {
    var x1 int
    t.Logf("name of x1: %v", reflect.TypeOf(x1).Name()) // int

    type MyInt int
    var x2 MyInt
    t.Logf("name of x2: %v", reflect.TypeOf(x2).Name()) // MyInt

    var y1 struct{}
    t.Logf("name of y1: %v", reflect.TypeOf(y1).Name()) // empty string

    type MyStruct struct{}
    var y2 MyStruct
    t.Logf("name of y2: %v", reflect.TypeOf(y2).Name()) // MyStruct

    var z1 []int
    t.Logf("name of z1: %v", reflect.TypeOf(z1).Name()) // empty string

    type MySlice []int
    var z2 MySlice
    t.Logf("name of z2: %v", reflect.TypeOf(z2).Name()) // MySlice
}
