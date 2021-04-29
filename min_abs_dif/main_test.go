package main

import "testing"

func TestAbs(t *testing.T) {
	var cases = []struct {
		i        int32
		expected int32
	}{
		{
			-80,
			80,
		},
		{
			100_000_000,
			100_000_000,
		},
		{
			-100_000_000,
			100_000_000,
		},
	}

	for _, c := range cases {
		a := Abs(c.i)
		if a != c.expected {
			t.Errorf("E: %d != %d", a, c.expected)
		}
	}
}

func TestMinAbsDiff(t *testing.T) {
	var cases = []struct {
		i        []int32
		expected int32
	}{
		{
			[]int32{3, -7, 0},
			3,
		},
		{
			[]int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75},
			1,
		},
		{
			[]int32{1, -3, 71, 68, 17},
			3,
		},
	}

	for _, c := range cases {
		res := minimumAbsoluteDifference(c.i)
		if res != c.expected {
			t.Errorf("E: %d != %d", res, c.expected)
		}
	}
}
