package builder

// IglooHouse is the struct that will be built
type NormalHouse struct {
	House
}

func newNormalHouse() *NormalHouse {
	return &NormalHouse{}
}

func (h *NormalHouse) SetWindows(windows int) {
	h.windows = windows
}

func (h *NormalHouse) SetDoors(doors int) {
	h.doors = doors
}

func (h *NormalHouse) SetFloors(floors int) {
	h.floors = floors
}

func (h *NormalHouse) GetWindows() int {
	return h.windows
}

func (h *NormalHouse) GetDoors() int {
	return h.doors
}

func (h *NormalHouse) GetFloors() int {
	return h.floors
}
func (h *NormalHouse) getHouse() House {
	return House{
		windows: h.GetWindows(),
		doors:   h.GetDoors(),
		floors:  h.GetFloors(),
	}
}
