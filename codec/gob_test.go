package codec

import (
    "bytes"
    "encoding/gob"
    "fmt"
    "io"
    "reflect"
    "testing"
)

type Person struct {
    Name string
    Age  int
}

func TestGobBasicEncodingAndDecoding(t *testing.T) {
    var buffer bytes.Buffer

    encoder := gob.NewEncoder(&buffer)
    decoder := gob.NewDecoder(&buffer)

    person := Person{Name: "Benjamin", Age: 42}

    err := encoder.Encode(person)
    if err != nil {
        t.Error("Encoding Error:", err)
        return
    }

    var decodedPerson Person
    err = decoder.Decode(&decodedPerson)
    if err != nil {
        t.Error("Decoding Error:", err)
        return
    }

    t.Log("Original Person:", person)
    t.Log("Decoded Person:", decodedPerson)
}

func TestGobEncodeMap(t *testing.T) {
    var buffer bytes.Buffer

    encoder := gob.NewEncoder(&buffer)
    value := reflect.ValueOf(map[string]int{"apple": 42, "banana": 89})
    err := encoder.EncodeValue(value)
    if err != nil {
        t.Error("Encoding Error:", err)
        return
    }

    decoder := gob.NewDecoder(&buffer)
    var decodedMap map[string]int
    err = decoder.Decode(&decodedMap)
    if err != nil {
        t.Error("Decoding Error:", err)
        return
    }

    t.Log("Original Map:", value)
    t.Log("Decoded Map:", decodedMap)
}

func encodePerson(writer io.Writer, person Person) error {
    encoder := gob.NewEncoder(writer)
    return encoder.Encode(person)
}

func decodePerson(reader io.Reader) (Person, error) {
    var person Person
    decoder := gob.NewDecoder(reader)
    err := decoder.Decode(&person)
    return person, err
}

func TestGobWithIOWriterAndReader(t *testing.T) {
    var buffer bytes.Buffer
    // bytes.Buffer implements both io.Writer and io.Reader by duck typing

    person := Person{Name: "Benjamin", Age: 89}

    err := encodePerson(&buffer, person)
    if err != nil {
        t.Error("Encoding Error:", err)
        return
    }

    fmt.Println("Encoded Gob Data:", buffer.Bytes())

    decodedPerson, err := decodePerson(&buffer)
    if err != nil {
        t.Error("Decoding Error:", err)
        return
    }

    t.Log("Original Person:", person)
    t.Log("Decoded Person:", decodedPerson)
}

func TestGobEncodeStruct(t *testing.T) {
    gob.Register(Person{})

    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    decoder := gob.NewDecoder(&buffer)

    original := Person{Name: "Alice", Age: 30}
    if err := encoder.Encode(original); err != nil {
        fmt.Println("Error encoding:", err)
        return
    }

    var decoded Person
    if err := decoder.Decode(&decoded); err != nil {
        fmt.Println("Error decoding:", err)
        return
    }

    fmt.Printf("Decoded data: %+v\n", decoded)
}

func TestGobEncodeMultipleObject(t *testing.T) {
    gob.Register(Person{})

    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    decoder := gob.NewDecoder(&buffer)

    original1 := Person{Name: "Alice", Age: 31}
    original2 := Person{Name: "Bob", Age: 32}

    if err := encoder.Encode(original1); err != nil {
        fmt.Println("Error encoding:", err)
        return
    }

    if err := encoder.Encode(original2); err != nil {
        fmt.Println("Error encoding:", err)
        return
    }

    var decoded1 Person
    if err := decoder.Decode(&decoded1); err != nil {
        fmt.Println("Error decoding:", err)
        return
    }

    var decoded2 Person
    if err := decoder.Decode(&decoded2); err != nil {
        fmt.Println("Error decoding:", err)
        return
    }

    fmt.Printf("Decoded data 1: %+v\n", decoded1)
    fmt.Printf("Decoded data 2: %+v\n", decoded2)
}

type Employee struct {
    Name   string
    salary int
}

func TestGobEncodeUnexportedField(t *testing.T) {
    gob.Register(Employee{})

    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    decoder := gob.NewDecoder(&buffer)

    original := Employee{Name: "Alice", salary: 6174}
    if err := encoder.Encode(original); err != nil {
        fmt.Println("Error encoding:", err)
        return
    }

    var decoded Employee
    if err := decoder.Decode(&decoded); err != nil {
        fmt.Println("Error decoding:", err)
        return
    }

    fmt.Printf("Decoded data: %+v\n", decoded)
    if decoded.salary != 0 {
        t.Error("Unexported field should not be encoded")
    }
}

func TestGobWithSpecialCase(t *testing.T) {
    gob.Register(Person{})

    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    decoder := gob.NewDecoder(&buffer)

    // decoded has non-default value
    // original has non-default value
    // decoded should be overwritten
    original := Person{Name: "Alice", Age: 30}
    if err := encoder.EncodeValue(reflect.ValueOf(original)); err != nil {
        fmt.Println("Error encoding:", err)
        return
    }

    var decoded Person
    decoded.Name = "Bob"
    if err := decoder.Decode(&decoded); err != nil {
        fmt.Println("Error decoding:", err)
        return
    }

    fmt.Printf("Decoded data: %+v\n", decoded)
    if decoded.Name != "Alice" {
        t.Error("Decoded data should not be overwritten")
    }

    // decoded has non-default value
    // original has default value
    // decoded should not be overwritten
    original.Name = ""
    if err := encoder.EncodeValue(reflect.ValueOf(original)); err != nil {
        fmt.Println("Error encoding:", err)
        return
    }

    decoded.Name = "Bob"
    if err := decoder.Decode(&decoded); err != nil {
        fmt.Println("Error decoding:", err)
        return
    }

    if decoded.Name != "Bob" {
        t.Error("Decoded data should be overwritten")
    }
}
