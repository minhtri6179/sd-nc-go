package main

import (
	"fmt"
	builder "root/Builder/builder"
)

func main() {
	normal := builder.GetBuilder("normal")
	director := builder.NewDirector(normal)
	normalHouse := director.BuildHouse()
	fmt.Println(normalHouse.GetDoors())
	fmt.Println(normalHouse.GetWindows())
	fmt.Println(normalHouse.GetFloors())

}
