package main

import "fmt"

type Node struct {
	Value int32
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type NodeQueue struct {
	Items []Node
}

func NewNodeQueue() *NodeQueue {
	return &NodeQueue{
		make([]Node, 100),
	}
}

func (q *NodeQueue) Enqueue(n *Node) {
	q.Items = append(q.Items, *n)
}

func (q *NodeQueue) Dequeue() *Node {
	n := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return &n
}

func (q *NodeQueue) Front() *Node {
	return &q.Items[0]
}

func (q *NodeQueue) Size() int {
	return len(q.Items)
}

func (q *NodeQueue) IsEmpty() bool {
	return len(q.Items) == 0
}

type Graph struct {
	Nodes []*Node
	Edges map[int32][]Node
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if _, ok := g.Edges[n1.Value]; !ok {
		g.Edges[n1.Value] = []Node{*n2}
	} else {
		g.Edges[n1.Value] = append(g.Edges[n1.Value], *n2)
	}
}

// ExistingNode manages duplicates
func (g *Graph) ExistingNode(n *Node) Node {
	for _, i := range g.Nodes {
		if i.Value == n.Value {
			return *i
		}
	}
	g.AddNode(n)
	return *n
}

func (g *Graph) BFS(idx int) (distances []int) {
	q := NewNodeQueue()
	n := g.Nodes[idx]
	q.Enqueue(n)
	visited := make(map[int32]bool)
	// -1 for orphan, +6 for (each intermediary) path
	distances = make([]int, len(g.Nodes))

	for {
		if q.IsEmpty() {
			break
		}

		node := q.Dequeue()
		visited[node.Value] = true
		neighbours := g.Edges[n.Value]

		for i := 0; i < len(neighbours); i++ {
			j := neighbours[i]
			if !visited[j.Value] {
				q.Enqueue(&j)
				visited[j.Value] = true
			}
		}

	}
	return distances
}
