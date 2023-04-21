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

func TestGobEncodeValue(t *testing.T) {
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
