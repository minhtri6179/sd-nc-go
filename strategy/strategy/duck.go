package strategy

// Duck is the main struct for the strategy pattern
type Duck struct {
	name        string
	age         int
	flyBehavior FlyBehavior
}

// NewDuck is the constructor for the Duck struct
func NewDuck(f FlyBehavior) *Duck {
	return &Duck{
		flyBehavior: f,
		name:        "Duck",
		age:         0,
	}
}

// PerformFly is the method that will be used to perform the fly behavior
func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

// PerformFly is the method that will be used to perform the fly behavior
func (d *Duck) PerformFly() {
	d.flyBehavior.fly(d)
}
