package main

import "testing"

func TestCreateTree(t *testing.T) {
	var cases = []struct {
		n        int
		indexes  [][]int32
		t        int32
		queries  []int32
		expected string
	}{
		{
			3,
			[][]int32{
				{2, 3},
				{-1, -1},
				{-1, -1},
			},
			int32(2),
			[]int32{1, 1},
			"",
		},
		{
			5,
			[][]int32{
				{2, 3},
				{-1, 4},
				{-1, 5},
				{-1, -1},
				{-1, -1},
			},
			int32(1),
			[]int32{2},
			"",
		},
	}

	for _, c := range cases {
		tree := CreateBinaryTree(c.indexes)
		inOrder(tree)
	}
}

func TestMirrorTree(t *testing.T) {
	var cases = []struct {
		n        int
		indexes  [][]int32
		t        int32
		queries  []int32
		expected string
	}{
		{
			3,
			[][]int32{
				{2, 3},
				{-1, -1},
				{-1, -1},
			},
			int32(2),
			[]int32{1, 1},
			"",
		},
		{
			5,
			[][]int32{
				{2, 3},
				{-1, 4},
				{-1, 5},
				{-1, -1},
				{-1, -1},
			},
			int32(1),
			[]int32{2},
			"",
		},
	}

	for _, c := range cases {
		tree := CreateBinaryTree(c.indexes)
		t.Logf("Before mirror\n")
		inOrder(tree)
		t.Logf("Mirror\n")
		mirror(tree)
		t.Logf("After Mirror\n")
		inOrder(tree)
	}
}

func TestMirrorTreeLevel(t *testing.T) {
	var cases = []struct {
		n        int
		indexes  [][]int32
		t        int32
		queries  []int32
		expected string
	}{
		{
			3,
			[][]int32{
				{2, 3},
				{-1, -1},
				{-1, -1},
			},
			int32(2),
			[]int32{1, 1},
			"",
		},
		{
			5,
			[][]int32{
				{2, 3},
				{-1, 4},
				{-1, 5},
				{-1, -1},
				{-1, -1},
			},
			int32(1),
			[]int32{2},
			"",
		},
	}

	for _, c := range cases {
		tree := CreateBinaryTree(c.indexes)
		t.Logf("Before Swap\n")
		inOrder(tree)
		for _, q := range c.queries {
			swapLevel(tree, q)
			t.Logf("\nAfter Swap\n")
			inOrder(tree)
		}
	}
}

func TestDecomposeTree(t *testing.T) {
	var cases = []struct {
		n        int
		indexes  [][]int32
		t        int32
		queries  []int32
		expected [][]int32
	}{
		{
			3,
			[][]int32{
				{2, 3},
				{-1, -1},
				{-1, -1},
			},
			int32(2),
			[]int32{1, 1},
			[][]int32{
				{2, 3},
				{-1, -1},
				{-1, -1},
			},
		},
		{
			5,
			[][]int32{
				{2, 3},
				{-1, 4},
				{-1, 5},
				{-1, -1},
				{-1, -1},
			},
			int32(1),
			[]int32{2},
			[][]int32{
				{2, 3},
				{-1, 4},
				{-1, 5},
				{-1, -1},
				{-1, -1},
			},
		},
	}

	for _, c := range cases {
		tree := CreateBinaryTree(c.indexes)
		t.Logf("Before Swap\n")
		inOrder(tree)
		res := decomposeBT(tree, len(c.indexes))

		if !compare(res, c.expected) {
			t.Errorf("E: Expected %v == %v", res, c.expected)
		}
	}
}

func TestDecomposeTreeInOrder(t *testing.T) {
	var cases = []struct {
		n        int
		indexes  [][]int32
		t        int32
		queries  []int32
		expected [][]int32
	}{
		{
			3,
			[][]int32{
				{2, 3},
				{-1, -1},
				{-1, -1},
			},
			int32(2),
			[]int32{1, 1},
			[][]int32{
				{2, 1, 3},
				{3, 1, 2},
			},
		},
		{
			5,
			[][]int32{
				{2, 3},
				{-1, 4},
				{-1, 5},
				{-1, -1},
				{-1, -1},
			},
			int32(1),
			[]int32{2},
			[][]int32{
				{4, 2, 1, 5, 3},
			},
		},
	}

	for _, c := range cases {
		d := make([][]int32, len(c.queries))
		tree := CreateBinaryTree(c.indexes)

		for j, q := range c.queries {
			decomposeIO(tree, d, j)
			swapLevel(tree, q)
		}

		if !compare(d, c.expected) {
			t.Errorf("E: Expected %v == %v", d, c.expected)
		}
	}
}

func compare(a, b [][]int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}
