package main

import (
	"container/ring"
	"fmt"
)

type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}

type RingBuffer struct {
	sendq *ring.Ring
	getq  *ring.Ring
}

func NewRingBuffer(size int) *RingBuffer {
	rb := &RingBuffer{}
	ring := ring.New(size)
	rb.sendq = ring
	rb.getq = ring
	return rb
}

func (rb *RingBuffer) Add(val int) {
	rb.sendq.Value = val
	if rb.sendq == rb.getq {
		rb.getq = rb.getq.Next()
	}
	rb.sendq = rb.sendq.Next()
}
func (rb *RingBuffer) Get() (int, bool) {
	val := rb.getq.Value
	if val == nil {
		return 0, false
	}
	rb.getq.Value = nil
	rb.getq = rb.getq.Next()
	return val.(int), true
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4)

	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 2
	}
	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 3
	}
	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 4
	}
	if _, ok := rb.Get(); !ok {
		fmt.Println("Буфер пуст") // Выводит: Буфер пуст
	}

}
