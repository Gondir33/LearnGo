package workers

import (
	"log"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	avl := GenerateTree(80)
	log.Println(avl.ToMermaid())
}
