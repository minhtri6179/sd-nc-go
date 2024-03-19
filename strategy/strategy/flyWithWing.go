package strategy

import "fmt"

type FlyWithWing struct {
}

func (f *FlyWithWing) fly(d *Duck) {
	fmt.Println("I'm flying with wings!")
}
