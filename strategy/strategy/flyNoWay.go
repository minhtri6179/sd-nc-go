package strategy

import "fmt"

type FlyNoWay struct {
}

func (f *FlyNoWay) fly(d *Duck) {
	fmt.Println("I can't fly!")
}
