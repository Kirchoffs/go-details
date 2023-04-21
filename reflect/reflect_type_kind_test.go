package reflect

import (
    "fmt"
    "reflect"
    "testing"
)

func TestMapTypeAndKind(t *testing.T) {
    var m map[string]int

    refObjTypOfM := reflect.TypeOf(m)
    refObjKindOfM := refObjTypOfM.Kind()
    t.Logf("Type of m: %v", refObjTypOfM)  // Type: map[string]int
    t.Logf("Kind of m: %v", refObjKindOfM) // Kind: map
}

func TestTypeUtility(t *testing.T) {
    CheckType := func(x interface{}) {
        typ := reflect.TypeOf(x)
        switch typ.Kind() {
        case reflect.Int:
            fmt.Println("x is an int")
            t.Logf("Type of x: %v", typ.String())
        case reflect.String:
            fmt.Println("x is a string")
            t.Logf("Type of x: %v", typ.String())
        case reflect.Slice:
            fmt.Println("x is a slice")
            t.Logf("Type of x: %v", typ.String())
        case reflect.Map:
            fmt.Println("x is a map")
            t.Logf("Type of x: %v", typ.String())
        default:
            fmt.Println("Unknown type")
            t.Logf("Type of x: %v", typ.String())
        }
    }

    CheckType(42)
    CheckType("hello")
    CheckType([]int{1, 2, 3})
    CheckType(map[string]int{"answer": 42})
}
