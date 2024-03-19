package strategy

// Duck is the main struct for the strategy pattern
type Duck struct {
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
}

// NewDuck is the constructor for the Duck struct
func NewDuck() *Duck {
	return &Duck{}
}

// PerformFly is the method that will be used to perform the fly behavior
func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.FlyBehavior = fb
}

// PerformFly is the method that will be used to perform the fly behavior
func (d *Duck) SetQuackBehavior(qb QuackBehavior) {
	d.QuackBehavior = qb
}

// PerformFly is the method that will be used to perform the fly behavior
func (d *Duck) PerformFly() string {
	return d.FlyBehavior.Fly()
}

// PerformQuack is the method that will be used to perform the quack behavior
func (d *Duck) PerformQuack() string {
	return d.QuackBehavior.Quack()
}
