package basic

import (
    "fmt"
    "testing"
)

type Quacker interface {
    Quack()
}

type Duck struct{}

func (d Duck) Quack() {
    fmt.Println("I can quack!")
}

type WoodenDuck struct{}

func (wd WoodenDuck) Quack() {
    fmt.Println("I can't really quack!")
}

type PointerDuck struct{}

func (pd *PointerDuck) Quack() {
    fmt.Println("I can quack as a pointer!")
}

func makeItQuack(q Quacker) {
    q.Quack()
}

func TestDuckTyping(t *testing.T) {
    duck := Duck{}
    woodenDuck := WoodenDuck{}
    pointerDuck := PointerDuck{}

    makeItQuack(duck)
    makeItQuack(woodenDuck)
    makeItQuack(&pointerDuck)
}
