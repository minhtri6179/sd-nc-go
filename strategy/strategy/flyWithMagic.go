package strategy

import "fmt"

// FlyWithMagic is a struct that will be used to implement the fly behavior
type FlyWithMagic struct {
}

func (f *FlyWithMagic) fly(d *Duck) {
	fmt.Println("I'm flying with magic!")
}
