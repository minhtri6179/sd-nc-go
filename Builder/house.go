package builder

// House is the struct that will be built
type House struct {
	windows int
	doors   int
	floors  int
}

func (h *House) SetWindows(windows int) {
	h.windows = windows
}

func (h *House) SetDoors(doors int) {
	h.doors = doors
}

func (h *House) SetFloors(floors int) {
	h.floors = floors
}

func (h *House) GetWindows() int {
	return h.windows
}

func (h *House) GetDoors() int {
	return h.doors
}

func (h *House) GetFloors() int {
	return h.floors
}
