package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Node struct {
	index int
	data  *User
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

// Зачем возвращать *BinaryTree?
func (t *BinaryTree) insert(user *User) *BinaryTree {
	if t.root == nil {
		t.root = &Node{data: user, index: 0}
		return t
	}

	t.root.insert(user)
	return t
}

func (n *Node) insert(user *User) {
	if user.ID == n.data.ID {
		fmt.Errorf("user with id = %v already exists\n", user.ID)
		return
	}

	if n.data.ID > user.ID {
		if n.left == nil {
			n.left = &Node{data: user, index: n.index + 1}
			return
		}
		n.left.insert(user)
	} else {
		if n.right == nil {
			n.right = &Node{data: user, index: n.index + 1}
			return
		}
		n.right.insert(user)
	}
}

func (t *BinaryTree) search(key int) *User {
	return t.root.search(key)
}

func (n *Node) search(key int) *User {
	if n == nil {
		return nil
	}
	if n.data.ID > key {
		return n.left.search(key)
	} else if n.data.ID < key {
		return n.right.search(key)
	} else {
		return n.data
	}
}

func generateData(n int) *BinaryTree {
	rand.Seed(time.Now().UnixNano())
	bt := &BinaryTree{}
	for i := 0; i < n; i++ {
		val := rand.Intn(100)
		bt.insert(&User{
			ID:   val,
			Name: fmt.Sprintf("User%d", val),
			Age:  rand.Intn(50) + 20,
		})
	}
	return bt
}

func main() {
	bt := generateData(50)
	user := bt.search(30)
	if user != nil {
		fmt.Printf("Найден пользователь: %+v\n", user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
