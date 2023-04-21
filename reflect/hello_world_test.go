package reflect

import (
    "reflect"
    "testing"
)

func TestHelloWorld(t *testing.T) {
    var x float64 = 3.14

    refObjValOfX := reflect.ValueOf(x)
    refObjTypOfX := reflect.TypeOf(x)                              // Or refObjTypOfX := refObjValOfX.Type()
    t.Logf("Reflect object value of x: %v", refObjValOfX.String()) // <float64 Value>
    t.Logf("Reflect object value of x: %f", refObjValOfX.Float())  // 3.14
    t.Logf("Reflect object type of x: %v", refObjTypOfX.String())  // float64
    t.Logf("Can set x: %v", refObjValOfX.CanSet())                 // false

    refObjValPtrOfX := reflect.ValueOf(&x)
    refObjTypPtrOfX := reflect.TypeOf(&x)                                    // Or refObjTypPtrOfX := refObjValPtrOfX.Type()
    t.Logf("Reflect object value of &x: %v", refObjValPtrOfX.String())       // <*float64 Value>
    t.Logf("Reflect object value of &x: %f", refObjValPtrOfX.Elem().Float()) // 3.14
    t.Logf("Reflect object type of &x: %v", refObjTypPtrOfX.String())        // *float64
    t.Logf("Can set &x: %v", refObjValPtrOfX.Elem().CanSet())                // true
}

func TestTypeAndKind(t *testing.T) {
    type Profile struct {
        name   string
        age    int
        gender string
    }

    profile := Profile{}
    refObjtypOfProfile := reflect.TypeOf(profile)
    t.Log("Type of profile: ", refObjtypOfProfile.String()) // Type: reflect.Profile
    t.Log("Kind of profile: ", refObjtypOfProfile.Kind())   // Kind: struct
}
