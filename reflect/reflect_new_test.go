package reflect

import (
    "reflect"
    "testing"
)

type MyStruct struct {
    Name string
    Age  int
}

func TestReflectNew(t *testing.T) {
    reflectType := reflect.TypeOf(MyStruct{})

    v := reflect.New(reflectType)
    elem := v.Elem()

    elem.FieldByName("Name").SetString("John Doe")
    elem.FieldByName("Age").SetInt(35)

    // elem.Field(0).SetString("John Doe")
    // elem.Field(1).SetInt(35)

    i := v.Interface()
    ms := i.(*MyStruct)

    t.Logf("Name: %s, Age: %d", ms.Name, ms.Age) // Name: John Doe, Age: 35
}
