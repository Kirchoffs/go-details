package basic

import "testing"

type Dog struct {
    name string
}

func (dog *Dog) setName(name string) {
    dog.name = name
}

func TestInterfaceAssignment(t *testing.T) {
    dog := Dog{"Tom"}

    var animal interface{} = dog

    dog.name = "Jerry"
    t.Logf("Dog name: %s", dog.name)             // Jerry
    t.Logf("Animal name: %s", animal.(Dog).name) // Tom

    dog.setName("Spike")
    t.Logf("Dog name: %s", dog.name)             // Spike
    t.Logf("Animal name: %s", animal.(Dog).name) // Tom

    var animalPointer interface{} = &dog
    dog.setName("Tyke")
    t.Logf("Dog name: %s", dog.name)                     // Tyke
    t.Logf("Animal name: %s", animalPointer.(*Dog).name) // Tyke
}
