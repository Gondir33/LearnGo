package main

import (
	"container/list"
	"fmt"
)

type Car struct {
	LicensePlate string
}

type ParkingLot struct {
	space *list.List
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{space: list.New()}
}

func (p *ParkingLot) Park(c Car) {
	fmt.Printf("Автомобиль {%v} припаркован\n", c.LicensePlate)
	p.space.PushBack(c)

}

func (p *ParkingLot) Leave() {
	tmp := p.space.Front()
	if tmp == nil {
		fmt.Println("Парковка пуста")
		return
	}
	fmt.Printf("Автомобиль %v покинул парковку\n", tmp.Value)
	p.space.Remove(tmp)
}

func main() {
	parkingLot := NewParkingLot()
	parkingLot.Park(Car{LicensePlate: "ABC-123"})
	parkingLot.Park(Car{LicensePlate: "XYZ-789"})
	parkingLot.Leave()
	parkingLot.Leave()
	parkingLot.Leave()
}
