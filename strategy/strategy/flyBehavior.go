package strategy

// Flyable is the interface that will be used to implement the fly behavior
type FlyBehavior interface {
	fly(duck *Duck)
}
