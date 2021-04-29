package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'roadsAndLibraries' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER c_lib
 *  3. INTEGER c_road
 *  4. 2D_INTEGER_ARRAY cities
 */

type Node struct {
	value int32
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	// lock sync.RWMutex
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	// g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *Graph) ExistingNode(n *Node) Node {
	for i := 0; i < len(g.nodes); i++ {
		if n.value == g.nodes[i].value {
			return *g.nodes[i]
		}
	}
	g.AddNode(n)
	return *n
}

func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + "-> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

type NodeQueue struct {
	items []Node
}

func (pq NodeQueue) New() *NodeQueue {
	return &NodeQueue{
		[]Node{},
	}
}

func (pq *NodeQueue) Enqueue(n Node) {
	pq.items = append(pq.items, n)
}

func (pq *NodeQueue) Dequeue() *Node {
	n := pq.items[0]
	pq.items = pq.items[1:len(pq.items)]
	return &n
}

func (pq *NodeQueue) Front() *Node {
	return &pq.items[0]
}

func (pq *NodeQueue) IsEmpty() bool {
	return len(pq.items) == 0
}

func (pq *NodeQueue) Size() int {
	return len(pq.items)
}

func (g *Graph) BFS(idx int, f func(*Node)) (connections int64) {
	q := NodeQueue{}
	q.New()
	// i by default should be zero
	n := g.nodes[idx]
	q.Enqueue(*n)
	visited := make(map[int32]bool)

	for {
		if q.IsEmpty() {
			break
		}

		node := q.Dequeue()
		visited[node.value] = true
		near := g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j.value] {
				q.Enqueue(*j)
				visited[j.value] = true
			}
		}
		if f != nil {
			f(node)
		}
	}
	connections = int64(0)
	for _, v := range visited {
		if v {
			connections++
		}
	}
	return
}

/*
Menger's Theorem
A simple algorithm might be written in pseudo-code as follows:

1. Begin at any arbitrary node of the graph, G
2. Proceed from that node using either depth-first or breadth-first search, counting all nodes reached.
3. Once the graph has been entirely traversed, if the number of nodes counted is equal to the number of nodes of G, the graph is connected; otherwise it is disconnected.

*May have to change map[int32]int64 back to map[Node]int64
*/
func (g *Graph) Connectivity() (connections map[int32]int64) {
	max := int64(-100_000_000_000)
	maxConn := len(g.nodes)
	connections = make(map[int32]int64)

	for i, n := range g.nodes {
		connections[n.value] = g.BFS(i, nil)
		if connections[n.value] == int64(maxConn) {
			// best possible path
			return connections
		} else if max < connections[n.value] {
			max = connections[n.value]
		}
	}
	return
}

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	// Write your code here
	g := &Graph{}
	for i := 0; i < len(cities); i++ {
		cityA := &Node{cities[i][0]}
		cityB := &Node{cities[i][1]}

		g.ExistingNode(cityA)
		g.ExistingNode(cityB)

		g.AddEdge(cityA, cityB)
	}

	connections := g.Connectivity()

	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
