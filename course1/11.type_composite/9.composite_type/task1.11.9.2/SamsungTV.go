package main

import "fmt"

type tv interface {
	switchOFF() bool
	switchOn() bool
	GetStatus() bool
	GetModel() string
	SamsungHub() string
	LgHub()
}
type LgTV struct {
	status bool
	model  string
}

type SamsungTV struct {
	status bool
	model  string
}

func (tv *SamsungTV) SamsungHub() string {
	tv.status = false
	return "SamsungHub"
}

func (tv *LgTV) LGHub() string {
	tv.status = false
	return "LGHub"
}

func (tv *SamsungTV) GetModel() string {
	return tv.model
}

func (tv *LgTV) GetModel() string {
	return tv.model
}

func (tv *SamsungTV) GetStatus() bool {
	return tv.status
}

func (tv *LgTV) GetStatus() bool {
	return tv.status
}

func (tv *SamsungTV) switchOn() bool {
	tv.status = true
	return true
}

func (tv *LgTV) switchOn() bool {
	tv.status = true
	return true
}

func (tv *SamsungTV) switchOFF() bool {
	tv.status = false
	return true
}

func (tv *LgTV) switchOFF() bool {
	tv.status = false
	return true
}

func main() {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.GetModel())   // Samsung XL-100500
	fmt.Println(tv.SamsungHub()) // SamsungHub
	fmt.Println(tv.GetStatus())  // false
	fmt.Println(tv.switchOn())   // true
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.switchOFF())  // true
	fmt.Println(tv.GetStatus())  // false
}
