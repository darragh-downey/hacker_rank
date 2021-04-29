package main

import (
	"fmt"
	"testing"
)

func TestRoadsNLibraries(t *testing.T) {
	var cases = []struct {
		n        int32
		cLib     int32
		cRoad    int32
		cities   [][]int32
		expected int64
	}{
		{
			3,
			2,
			1,
			[][]int32{{1, 2}, {3, 1}, {2, 3}},
			4,
		},
		{
			6,
			2,
			5,
			[][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}},
			12,
		},
	}

	for _, c := range cases {
		res := roadsAndLibraries(c.n, c.cLib, c.cRoad, c.cities)
		if res != c.expected {
			t.Errorf("E: %d != %d\n", res, c.expected)
		}
	}
}

func TestBFS(t *testing.T) {
	var cases = []struct {
		n        int32
		cLib     int32
		cRoad    int32
		cities   [][]int32
		expected int64
	}{
		{
			3,
			2,
			1,
			[][]int32{{1, 2}, {3, 1}, {2, 3}},
			4,
		},
		{
			6,
			2,
			5,
			[][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}},
			12,
		},
	}

	for _, c := range cases {
		g := &Graph{}
		for i := 0; i < len(c.cities); i++ {
			cityA := &Node{c.cities[i][0]}
			cityB := &Node{c.cities[i][1]}

			g.ExistingNode(cityA)
			g.ExistingNode(cityB)

			g.AddEdge(cityA, cityB)
		}

		g.BFS(0, func(n *Node) {
			fmt.Printf("%v\n", n)
		})
		fmt.Println()
	}
}

func TestConnectivity(t *testing.T) {
	var cases = []struct {
		n        int32
		cLib     int32
		cRoad    int32
		cities   [][]int32
		expected int64
	}{
		{
			3,
			2,
			1,
			[][]int32{{1, 2}, {3, 1}, {2, 3}},
			3,
		},
		{
			6,
			2,
			5,
			[][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}},
			6,
		},
	}

	for _, c := range cases {
		g := &Graph{}
		for i := 0; i < len(c.cities); i++ {
			cityA := &Node{c.cities[i][0]}
			cityB := &Node{c.cities[i][1]}

			g.ExistingNode(cityA)
			g.ExistingNode(cityB)

			g.AddEdge(cityA, cityB)
		}

		for node, connections := range g.Connectivity() {
			if connections == c.expected {
				t.Logf("I: Node %v has %d connections which is full connectivity", node, connections)
			} else {
				t.Logf("I: Node %v has %d connections", node, connections)
			}
		}

		fmt.Println()
	}
}
