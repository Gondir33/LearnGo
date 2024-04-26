package workers

import (
	"log"
	"testing"
)

func TestGraph(t *testing.T) {
	nodes := CreateRandomNodes(5, 30)
	buffer := BufferMermaidChart(nodes)
	log.Println(buffer.String())
}
