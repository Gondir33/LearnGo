package main

import "testing"

func TestGetStatus(t *testing.T) {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	res := tv.GetStatus() // true
	exp := true
	if res != exp {
		t.Errorf("get status don't work")
	}
}

func TestGetModel(t *testing.T) {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	res := tv.GetModel()
	exp := "Samsung XL-100500"
	if res != exp {
		t.Errorf("get model don't work")
	}
}

func TestSamsungHUB(t *testing.T) {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	res := tv.SamsungHub()
	exp := "SamsungHub"
	if res != exp && tv.status != false {
		t.Errorf("SamsunHUB don't work")
	}
}

func TestSwitchOn(t *testing.T) {
	tv := &SamsungTV{
		status: false,
		model:  "Samsung XL-100500",
	}
	res := tv.switchOn()
	exp := true
	if res != exp && tv.status != true {
		t.Errorf("SwitchOn don't work")
	}
}

func TestSwitchOFF(t *testing.T) {
	tv := &SamsungTV{
		status: false,
		model:  "Samsung XL-100500",
	}
	res := tv.switchOFF()
	exp := true
	if res != exp && tv.status != false {
		t.Errorf("SwitchOn don't work")
	}
}

func TestLGHUB(t *testing.T) {
	tv := &LgTV{
		status: true,
		model:  "Lg XL-100500",
	}
	res := tv.LGHub()
	exp := "LGHub"
	if res != exp && tv.status != false {
		t.Errorf("SamsunHUB don't work")
	}
}
