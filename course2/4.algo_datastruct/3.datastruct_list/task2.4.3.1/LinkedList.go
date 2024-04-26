package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n int) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

func partition(low, high *Node) *Node {
	pivot := high.data

	i := low.prev

	for j := low; j != high; j = j.next {
		if j.data.Date.UnixNano() <= pivot.Date.UnixNano() {
			if i == nil {
				i = low
			} else {
				i = i.next
			}
			i.data, j.data = j.data, i.data
		}
	}
	if i == nil {
		i = low
	} else {
		i = i.next
	}
	i.data, high.data = high.data, i.data
	return i
}

func quickSort(low, high *Node) {
	if high != nil && low != high && low != high.next {
		pivotIndex := partition(low, high)
		quickSort(low, pivotIndex.prev)
		quickSort(pivotIndex.next, high)
	}
}

// LoadData загрузка данных из подготовленного json файла
func (d *DoubleLinkedList) LoadData(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var commits []Commit
	err = json.Unmarshal(bytes, &commits)
	if err != nil {
		return err
	}
	// отсортировать список используя самописный QuickSort
	d.Init(commits)
	quickSort(d.head, d.tail)
	d.curr = d.head
	return nil
}

// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
	d.curr = d.curr.next
	return d.curr
}

// Prev получение предыдущего элемента
func (d *DoubleLinkedList) Prev() *Node {
	d.curr = d.curr.prev
	return d.curr
}

// Insert вставка элемента после n элемента
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n >= d.len || n < 0 {
		return errors.New("n is out of range")
	}
	curr := d.head
	for i := 0; i < n; i++ {
		curr = curr.next
	}
	next := curr.next

	if next == nil {
		curr.next = &Node{data: &c, prev: curr, next: next}
		d.tail = curr.next
	} else {
		curr.next = &Node{data: &c, prev: curr, next: next}
		next.prev = curr.next
	}
	d.len++
	return nil
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {
	if n >= d.len || n < 0 {
		return errors.New("n is out of range")
	}

	curr := d.head
	var repl *Node
	var flagToReplace bool
	for i := 0; i < n; i++ {
		curr = curr.next
	}
	if d.curr == curr {
		if d.curr.next != nil {
			repl = d.curr.next
			flagToReplace = true
		}
		if d.curr.prev != nil {
			repl = d.curr.prev
			flagToReplace = true
		}
	}
	if curr.next == nil {
		d.tail = curr.prev
		curr.prev.next = nil
	} else if curr == d.head {
		d.head = curr.next
		curr.next.prev = nil
	} else {
		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	}
	if flagToReplace == true {
		d.curr = repl
	}
	d.len--
	return nil
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.len == 0 {
		return errors.New("DoubleLinkedList is already empty")
	}
	curr := d.curr
	var repl *Node
	if d.curr.next != nil {
		repl = d.curr.next
	}
	if d.curr.prev != nil {
		repl = d.curr.prev
	}
	if curr.next == nil {
		d.tail = curr.prev
		curr.prev.next = nil
	} else if curr == d.head {
		d.head = curr.next
		curr.next.prev = nil
	} else {
		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	}
	d.curr = repl
	d.len--
	return nil
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	if d.len == 0 {
		return 0, errors.New("DoubleLinkedList is already empty")
	}
	var res int
	i := d.head
	for d.curr != i {
		res++
		i = i.next
	}
	return res, nil
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	return d.tail
}

// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
	return d.head
}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	curr := d.head
	for curr != nil {
		if curr.data.UUID == uuID {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// Search поиск коммита по message
func (d *DoubleLinkedList) Search(message string) *Node {
	curr := d.head
	for curr != nil {
		if curr.data.Message == message {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	newHead := &Node{data: d.tail.data, prev: nil, next: d.tail.prev}
	dReverse := &DoubleLinkedList{head: newHead, curr: newHead, len: d.len}
	curr := d.head
	for i := 0; i < d.len-1; i++ {
		newNext := &Node{data: curr.data, prev: dReverse.curr, next: nil}
		dReverse.curr.next = newNext
		curr = curr.next
		dReverse.Next()
	}
	dReverse.tail = &Node{data: d.head.data, prev: dReverse.curr.prev, next: nil}
	return dReverse
}

func (d *DoubleLinkedList) Init(c []Commit) {
	prev := &Node{data: &c[0], prev: nil, next: nil}
	d.head = prev
	d.curr = prev
	d.len = len(c)
	for i := 1; i < len(c); i++ {
		curr := &Node{data: &c[i], prev: prev, next: nil}
		prev.next = curr
		prev = prev.next
	}
	d.tail = prev
}

func (d *DoubleLinkedList) SetCurrent(n int) error {
	if n < 0 || n >= d.len {
		return errors.New("n is out of range")
	}

	curr := d.head
	for i := 0; i < n; i++ {
		curr = curr.next
	}
	d.curr = curr
	return nil
}

func (d *DoubleLinkedList) Push(c Commit) error {
	node := &Node{data: &c, next: nil, prev: d.tail}
	d.tail.next = node
	d.tail = node
	d.len++
	return nil
}

func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	if n < 0 || n >= d.len {
		return &Node{}, errors.New("n is out of range")
	}
	curr := d.head
	for i := 0; i < n; i++ {
		curr = curr.next
	}
	return curr, nil
}

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type Commit struct {
	Message string    `json:"message"`
	UUID    string    `json:"uuid"`
	Date    time.Time `json:"date"`
}

func GenerateData() []Commit {
	res := make([]Commit, 0, 10)
	for i := 0; i < 10; i++ {
		res = append(res, Commit{
			Message: gofakeit.Noun() + " " + gofakeit.Verb(),
			UUID:    gofakeit.UUID(),
			Date:    gofakeit.Date(),
		})
	}
	return res
}
