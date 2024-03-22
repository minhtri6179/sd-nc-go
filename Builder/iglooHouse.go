package builder

// IglooHouse is the struct that will be built

type IglooHouse struct {
	House
}

func newIglooHouse() *IglooHouse {
	return &IglooHouse{}
}

func (h *IglooHouse) SetWindows(windows int) {
	h.windows = windows
}

func (h *IglooHouse) SetDoors(doors int) {
	h.doors = doors
}

func (h *IglooHouse) SetFloors(floors int) {
	h.floors = floors
}

func (h *IglooHouse) getHouse() House {
	return House{
		windows: h.GetWindows(),
		doors:   h.GetDoors(),
		floors:  h.GetFloors(),
	}
}
