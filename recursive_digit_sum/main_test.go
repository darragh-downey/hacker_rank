package main

import (
	"strconv"
	"testing"
)

func TestSuperDigit(t *testing.T) {
	var cases = []struct {
		n        string
		k        int32
		expected string
	}{
		{
			"148",
			3,
			"3",
		},
		{
			"9875",
			4,
			"8",
		},
		{
			"861568688536788",
			100000,
			"3",
		},
	}

	for _, c := range cases {
		res := strconv.FormatInt(int64(superDigit(c.n, c.k)), 10)
		if res != c.expected {
			t.Errorf("E: %s != %s", res, c.expected)
		}
	}
}
