package main

import (
	"fmt"
	s "root/Strategy/strategy"
)

func main() {
	flyWithWing := &s.FlyWithWing{}
	flyNoWay := &s.FlyNoWay{}
	flyWithMagic := &s.FlyWithMagic{}
	duck := s.NewDuck(flyWithWing)
	duck.PerformFly()
	duck.SetFlyBehavior(flyNoWay)
	duck.PerformFly()
	duck.SetFlyBehavior(flyWithMagic)
	duck.PerformFly()
	fmt.Println("=====================================")
	eatMeat := &s.EatMeat{}
	eatGrass := &s.EatVegetable{}
	duck.SetEatBehavior(eatMeat)
	duck.PerformEat()
	duck.SetEatBehavior(eatGrass)
	duck.PerformEat()

}
