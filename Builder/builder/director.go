package builder

// House is the struct that will be built
type Director struct {
	builder IBBuilder
}

func NewDirector(builder IBBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) setBuilder(builder IBBuilder) {
	d.builder = builder
}

func (d *Director) BuildHouse() House {
	d.builder.SetWindows(4)
	d.builder.SetDoors(2)
	d.builder.SetFloors(1)
	return *d.builder.GetHouse()
}
