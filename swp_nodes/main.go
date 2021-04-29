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
 * A binary tree has the following characteristics:
 * It can be empty (null)
 * It contains a root node only
 * It contains a root node with a left subtree, a right subtree or both. These subtrees are also binary trees.
 */

type Tree struct {
	Value int32
	Level int32
	Left  *Tree
	Right *Tree
}

func NewTree(val, lvl int32) *Tree {
	return &Tree{
		Value: val,
		Level: lvl,
		Left:  nil,
		Right: nil,
	}
}

func (t *Tree) AddBranch(left, right *Tree) {
	t.Left = left
	t.Right = right
}

/*

Function Description

Complete the swapNodes function in the editor below. It should return a two-dimensional array where each element is an array of integers representing the node indices of an in-order traversal after a swap operation.

swapNodes has the following parameter(s):
- indexes: an array of integers representing index values of each
, beginning with , the first element, as the root.
- queries: an array of integers, each representing a k value.

Input Format

The first line contains n, number of nodes in the tree.

Each of the next n lines contains two integers, a b, where a is the index of left child, and b is the index of right child of ith node.

Note: -1 is used to represent a null node.

The next line contains an integer, t, the size of queries.
Each of the next t lines contains an integer queries[i], each being a value k.

Output Format

For each k, perform the swap operation and store the indices of your in-order traversal to your result array. After all swap operations have been performed, return your result array for printing.

Constraints

1 <= n <= 1024
1 <= t <= 100
1 <= k <= n
Either a = -1 or 2 <= a <= n
Either b = -1 or 2 <= b <= n
The index of a non-null child will always be greater than that of its parent
*/
func swapNodes(indexes [][]int32, queries []int32) (d [][]int32) {
	/*
	 * Write your code here.
	 */
	d = make([][]int32, len(queries))
	t := CreateBinaryTree(indexes)

	for i, q := range queries {
		swapLevel(t, q)
		decomposeIO(t, d, i)
	}

	return
}

/*
Each of the levels in indexes contains two integers, a b, where
a is the index of left child in indexes, and
b is the index of right child of ith/current node in indexes.
Note: -1 is used to represent a null node.
*/
func CreateBinaryTree(indexes [][]int32) *Tree {
	return createBT(int32(1), int32(1), int32(0), indexes)
}

func createBT(val, lvl, idx int32, indexes [][]int32) *Tree {
	if val < 0 {
		return nil
	}
	t := NewTree(val, lvl)

	left := createBT(indexes[idx][0], lvl+1, indexes[idx][0]-1, indexes)  // left node
	right := createBT(indexes[idx][1], lvl+1, indexes[idx][1]-1, indexes) // right node

	t.AddBranch(left, right)
	return t
}

// inverse of createBT; transform Binary Tree to [][]int32
func decomposeBT(t *Tree, l int) (d [][]int32) {
	d = make([][]int32, l)

	decompose(t, d)

	return d
}

func decompose(t *Tree, d [][]int32) {
	if t == nil {
		return
	}

	l, r := int32(-1), int32(-1)

	decompose(t.Left, d)
	decompose(t.Right, d)

	if t.Left != nil {
		l = t.Left.Value
	}
	if t.Right != nil {
		r = t.Right.Value
	}

	d[t.Value-1] = []int32{l, r}
}

func decomposeIO(t *Tree, d [][]int32, idx int) {
	if t == nil {
		return
	}

	decomposeIO(t.Left, d, idx)
	d[idx] = append(d[idx], t.Value)
	decomposeIO(t.Right, d, idx)
}

// Traverse left-subtree, visit root, traverse right-subtree
func inOrder(t *Tree) {
	if t == nil {
		return
	}

	inOrder(t.Left)
	fmt.Printf("%d ", t.Value)
	inOrder(t.Right)
}

// see 9 mirror() solution - http://cslibrary.stanford.edu/110/BinaryTrees.html
func mirror(t *Tree) {
	if t == nil {
		return
	}

	mirror(t.Left)
	mirror(t.Right)

	t.Left, t.Right = t.Right, t.Left
}

/*
Swap operation:

We define depth of a node as follows:

    The root node is at depth 1.
    If the depth of the parent node is d, then the depth of current node will be d+1.

Given a tree and an integer, k, in one operation, we need to swap the subtrees of all the nodes at each depth h, where h ∈ [k, 2k, 3k,...]. In other words, if h is a multiple of k, swap the left and right subtrees of that level.

You are given a tree of n nodes where nodes are indexed from [1..n] and it is rooted at 1. You have to perform t swap operations on it, and after each swap operation print the in-order traversal of the current state of the tree.
*/
func swapLevel(t *Tree, lvl int32) {
	if t == nil {
		return
	}

	swapLevel(t.Left, lvl)
	swapLevel(t.Right, lvl)

	if t.Level%lvl == 0 {
		t.Left, t.Right = t.Right, t.Left
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("/home/ddowney/Workspace/github.com/hackerrank/swap_nodes/output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var indexes [][]int32
	for indexesRowItr := 0; indexesRowItr < int(n); indexesRowItr++ {
		indexesRowTemp := strings.Split(readLine(reader), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != int(2) {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int32

	for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for resultRowItr, rowItem := range result {
		for resultColumnItr, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if resultColumnItr != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if resultRowItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
