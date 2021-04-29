package main

import "testing"

func TestBFS(t *testing.T) {
	var cases = []struct {
		in       [][]int32
		idx      int
		expected []int
	}{
		{
			[][]int32{{4, 2}, {1, 2}, {1, 3}},
			1,
			[]int{6, 6, -1},
		},
		{
			[][]int32{{3, 1}, {2, 3}},
			2,
			[]int{-1, 6},
		},
	}

	for _, c := range cases {
		g := &Graph{
			Nodes: make([]*Node, len(c.in)),
			Edges: make(map[int32][]Node),
		}
		for _, i := range c.in {
			n1 := g.ExistingNode(&Node{i[0]})
			n2 := g.ExistingNode(&Node{i[1]})
			g.AddEdge(&n1, &n2)
		}
		distances := g.BFS(c.idx)

		if !cmp(distances, c.expected) {
			t.Errorf("E: %v != %v", distances, c.expected)
		}
	}
}

func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
