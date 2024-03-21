package strategy

import "fmt"

type EatMeat struct {
}

func (e *EatMeat) eat(duck *Duck) {
	fmt.Println("I eat meat!")
}
