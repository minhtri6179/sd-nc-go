package builder

// House is the struct that will be built
type IBBuilder interface {
	SetWindows(windows int)
	SetDoors(doors int)
	SetFloors(floors int)
	GetWindows() int
	GetDoors() int
	GetFloors() int
}

func GetBuilder(builderType string) IBBuilder {
	switch builderType {
	case "normal":
		return &NormalHouse{}
	case "igloo":
		return &IglooHouse{}
	default:
		return &NormalHouse{}
	}
}
