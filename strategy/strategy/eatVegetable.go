package strategy

import "fmt"

// EatVegetable is the struct that will be used to implement the eat behavior
type EatVegetable struct {
}

// eat is the method that will be used to perform the eat behavior
func (e *EatVegetable) eat(duck *Duck) {
	fmt.Println("I eat vegetables!")
}

// NewEatVegetable is the constructor for the EatVegetable struct
func NewEatVegetable() *EatVegetable {
	return &EatVegetable{}
}
