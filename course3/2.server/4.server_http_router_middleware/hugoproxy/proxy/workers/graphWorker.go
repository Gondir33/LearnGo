package workers

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type GraphNode struct {
	ID    int
	Name  string // made random beer name
	Form  string // "circle"(()), "rect"[], "square"[], "ellipse"([]), "round-rect"(), "rhombus"{}
	Links []*GraphNode
}

func RandomForm() string {
	Forms := []string{"circle", "rect", "square", "ellipse", "round-rect", "rhombus"}

	index := rand.Intn(len(Forms))

	return Forms[index]
}

// should it be random???
// keep it simple and stupid?
// i --> i + 1
// i --> i + 2
// 0-->1
// 0-->2
// 1-->2
// 1-->3
func CreateLinks(nodes []*GraphNode) {
	for i := 0; i < len(nodes)-2; i++ {
		nodes[i].Links = append(nodes[i].Links, nodes[i+1])
		nodes[i].Links = append(nodes[i].Links, nodes[i+2])
	}
	nodes[len(nodes)-2].Links = append(nodes[len(nodes)-2].Links, nodes[len(nodes)-1])
}

func CreateRandomNodes(minSize, maxSize int) []*GraphNode {
	sz := gofakeit.IntRange(minSize, maxSize)

	nodes := make([]*GraphNode, sz)

	for i := 0; i < len(nodes); i++ {
		nodes[i] = &GraphNode{
			ID:    i,
			Form:  RandomForm(),
			Name:  gofakeit.BeerName(),
			Links: make([]*GraphNode, 0, 2),
		}
	}

	CreateLinks(nodes)

	return nodes
}

/*
<--->

{{< mermaid >}}
graph LR
A[Square Rect] --> B((Circle))
A --> C(Round Rect)
B --> D{Rhombus}
C --> D
C --> B
{{< /mermaid >}}
*/

const headGraph = `---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

`

func BufferMermaidChart(nodes []*GraphNode) *bytes.Buffer {
	var buffer bytes.Buffer

	formsMap := NewFormMap()
	elems := make([]string, 2)

	buffer.WriteString(headGraph)
	buffer.WriteString("{{< mermaid >}}\n")
	buffer.WriteString("graph LR\n")
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes[i].Links); j++ {
			elems[0] = strconv.Itoa(nodes[i].ID) +
				formsMap[nodes[i].Form][0] +
				nodes[i].Name +
				formsMap[nodes[i].Form][1]

			elems[1] = strconv.Itoa(nodes[i].Links[j].ID) +
				formsMap[nodes[i].Links[j].Form][0] +
				nodes[i].Links[j].Name +
				formsMap[nodes[i].Links[j].Form][1] + "\n"

			buffer.WriteString(strings.Join(elems, " --> "))
		}
	}

	buffer.WriteString("{{< /mermaid >}}\n")
	return &buffer
}

// "circle"(()), "rect"[], "square"[], "ellipse"([]), "round-rect"(), "rhombus"{}
func NewFormMap() map[string][]string {
	formsMap := make(map[string][]string, 6)

	formsMap["circle"] = []string{"((", "))"}
	formsMap["rect"] = []string{"[", "]"}
	formsMap["square"] = []string{"[", "]"}
	formsMap["ellipse"] = []string{"([", "])"}
	formsMap["round-rect"] = []string{"(", ")"}
	formsMap["rhombus"] = []string{"{", "}"}

	return formsMap
}

func GraphWorker() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			nodes := CreateRandomNodes(5, 30)
			buffer := BufferMermaidChart(nodes)
			err := os.WriteFile("/app/static/tasks/graph.md", buffer.Bytes(), 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
