package workers

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type TreeNode struct {
	Key    int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

type AVLTree struct {
	Root *TreeNode
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

const headAVLTree = `---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---


# Задача построить сбалансированное бинарное дерево

`

func WriteTree(node *TreeNode, buffer *bytes.Buffer) {
	if node == nil {
		return
	}
	if node.Left != nil {
		arg1 := strconv.Itoa(node.Key)
		arg2 := strconv.Itoa(node.Left.Key)
		buffer.WriteString(arg1 + " --> " + arg2 + "\n")
	}
	if node.Right != nil {
		arg1 := strconv.Itoa(node.Key)
		arg2 := strconv.Itoa(node.Right.Key)
		buffer.WriteString(arg1 + " --> " + arg2 + "\n")
	}
	WriteTree(node.Left, buffer)
	WriteTree(node.Right, buffer)

}

func EachNodeBalance(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	node.Left = EachNodeBalance(node.Left)
	node.Right = EachNodeBalance(node.Right)
	return balance(node)
}

func (t *AVLTree) ToMermaid() *bytes.Buffer {
	var buffer bytes.Buffer

	t.Root = EachNodeBalance(t.Root)

	buffer.WriteString(headAVLTree)
	buffer.WriteString("{{< mermaid >}}\n")
	buffer.WriteString("graph TD\n")

	WriteTree(t.Root, &buffer)

	buffer.WriteString("{{< /mermaid >}}\n")
	return &buffer
}

func NewNode(key int) *TreeNode {
	return &TreeNode{Key: key, Height: 1}
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func rightRotate(y *TreeNode) *TreeNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

func leftRotate(x *TreeNode) *TreeNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

func insert(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		return node
	}

	node.Height = max(height(node.Left), height(node.Right)) + 1

	return balance(node)
}

func balance(node *TreeNode) *TreeNode {
	balanceFactor := getBalance(node)

	if balanceFactor == 2 {
		if getBalance(node.Left) == -1 {
			node.Left = leftRotate(node.Left)
		}
		return rightRotate(node)

	}

	if balanceFactor == -2 {
		if getBalance(node.Right) == 1 {
			node.Right = rightRotate(node.Right)
		}
		return leftRotate(node)
	}

	return node
}

func GenerateTree(count int) *AVLTree {
	avl := &AVLTree{}
	for count > 0 {
		avl.Insert(rand.Intn(100))
		count--
	}

	return avl
}

func BinaryTreeWorker() {
	t := time.NewTicker(5 * time.Second)
	var avl *AVLTree
	count := 5
	for {
		select {
		case <-t.C:
			if count == 100 || count == 5 {
				avl = GenerateTree(count)
			} else {
				avl.Insert(rand.Intn(100))
			}
			err := os.WriteFile("/app/static/tasks/binary.md", avl.ToMermaid().Bytes(), 0644)
			if err != nil {
				log.Println(err)
			}
			count++
		}
	}
}
