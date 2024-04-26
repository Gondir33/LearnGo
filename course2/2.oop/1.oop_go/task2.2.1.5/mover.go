package main

import (
	"fmt"
	"strconv"
)

const (
	maxspeed = 120
	minspeed = 10
)

type Mover interface {
	Move() string
	Speed() int
	MaxSpeed() int
	MinSpeed() int
}

type BaseMover struct {
	speed int
}

func (b BaseMover) Move() string {
	return "Base mover. Moving at speed: " + strconv.Itoa(b.speed)
}

func (b BaseMover) Speed() int {
	return b.speed
}

func (b BaseMover) MaxSpeed() int {
	return maxspeed
}

func (b BaseMover) MinSpeed() int {
	return minspeed
}

type FastMover struct {
	BaseMover
}

func (f FastMover) Move() string {
	return "Fast mover! Moving at speed: " + strconv.Itoa(f.speed)
}

// func (f FastMover) Speed() int {
// 	return f.speed
// }

// func (f FastMover) MaxSpeed() int {
// 	return maxspeed
// }

// func (f FastMover) MinSpeed() int {
// 	return minspeed
// }

type SlowMover struct {
	BaseMover
}

func (s SlowMover) Move() string {
	return "Slow mover... Moving at speed: " + strconv.Itoa(s.speed)
}

// func (s SlowMover) Speed() int {
// 	return s.speed
// }

// func (s SlowMover) MaxSpeed() int {
// 	return maxspeed
// }

// func (s SlowMover) MinSpeed() int {
// 	return minspeed
// }

func main() {
	var movers []Mover
	fm := FastMover{BaseMover{100}}
	sm := SlowMover{BaseMover{10}}
	movers = append(movers, fm, sm)

	for _, mover := range movers {
		fmt.Println(mover.Move())
		fmt.Println("Maximum speed:", mover.MaxSpeed())
		fmt.Println("Minimum speed:", mover.MinSpeed())
	}
}
