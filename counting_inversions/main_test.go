package main

import "testing"

func TestCountInversions(t *testing.T) {
	var cases = []struct {
		i        []int32
		expected int
	}{
		{[]int32{1, 3, 5, 7}, 0},
		{[]int32{3, 2, 1}, 3},
		{[]int32{1, 1, 1, 2, 2}, 0},
		{[]int32{2, 1, 3, 1, 2}, 4},
		{[]int32{1, 5, 3, 7}, 1},
		{[]int32{7, 5, 3, 1}, 6},
	}

	for _, c := range cases {
		res := countInversions(c.i, int64(c.expected))
		if res != int64(c.expected) {
			t.Errorf("%d != %d", res, c.expected)
		}
	}
}
