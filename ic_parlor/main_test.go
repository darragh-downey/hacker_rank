package main

import "testing"

func TestWhatFlavours(t *testing.T) {
	var cases = []struct {
		money    int32
		cost     []int32
		expected string
	}{
		{
			int32(4),
			[]int32{1, 4, 5, 3, 2},
			"1 4",
		},
		{
			int32(4),
			[]int32{2, 2, 4, 3},
			"1 2",
		},
		{
			int32(5),
			[]int32{1, 2, 3, 5, 6},
			"2 3",
		},
		{
			int32(8),
			[]int32{4, 3, 2, 5, 7},
			"2 4",
		},
		{
			int32(12),
			[]int32{7, 2, 5, 4, 11},
			"1 3",
		},
	}

	for idx, c := range cases {
		res := whatFlavors(c.cost, c.money)
		if res != c.expected {
			t.Errorf("\nE: Failed case %d with:\n%v != %v", idx, res, c.expected)
		}
	}
}
