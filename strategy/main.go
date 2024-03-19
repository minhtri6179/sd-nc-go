package main

import s "root/Strategy/strategy"

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
}
