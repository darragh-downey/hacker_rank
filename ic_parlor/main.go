package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Set []int

func NewSet(v int) *Set {
	return &Set{v}
}

func (s Set) Add(v int) Set {
	for _, i := range s {
		if v == i {
			return s
		}
	}
	s = append(s, v)
	return s
}

func (s Set) Exists(v int) bool {
	for _, i := range s {
		if v == i {
			return true
		}
	}
	return false
}

func (s Set) Delete(v int) Set {
	idx := -1
	for i, k := range s {
		if k == v {
			idx = i
			break
		}
	}

	if idx != -1 {
		copy(s[idx:], s[idx+1:])
		s[len(s)-1] = -1
		s = s[:len(s)-1]
	}
	return s
}

// see approach 3 - https://leetcode.com/problems/two-sum/solution/
func whatFlavorsTwo(cost []int32, money int32) {
	idx := make(map[int32]Set)

	for i, v := range cost {
		complement := money - v
		if _, ok := idx[complement]; ok {
			fmt.Printf("%d %d\n", idx[complement][0], i)
			return
		}
		idx[v] = idx[v].Add(i)
	}
}

// Complete the whatFlavors function below.
func whatFlavors(cost []int32, money int32) string {
	idx := make([]int, 0, 2)

	costs := buildhash(money, cost)

	for i := 0; i < len(cost); i++ {
		if cost[i] >= money {
			continue
		}
		rem := money - cost[i]

		if _, ok := costs[rem]; ok {
			idx = append(idx, i)
			if len(costs[rem]) > 1 {
				for _, j := range costs[rem] {
					if j != i {
						idx = append(idx, j)
						break
					}
				}
			} else {
				idx = append(idx, costs[rem][0])
			}
			break
		}
	}

	return fmt.Sprintf("%d %d", idx[0]+1, idx[1]+1)
}

func buildhash(money int32, cost []int32) (costs map[int32][]int) {
	costs = make(map[int32][]int)

	for i := 0; i < len(cost); i++ {
		rem := money - cost[i]
		if rem <= 0 {
			continue
		}
		for j := 0; j < len(cost); j++ {
			if rem == cost[j] && i != j {
				if _, ok := costs[rem]; !ok {
					costs[rem] = []int{j}
				} else {
					costs[rem] = append(costs[rem], j)
				}
			}
		}
	}
	return
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	f, err := os.Open("/home/ddowney/Workspace/github.com/hackerrank/ice_cream_parlor/input13.txt")
	if err != nil {
		fmt.Printf("E: Failed to open file %s", f.Name())
		return
	}
	reader := bufio.NewReaderSize(f, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavorsTwo(cost, money)
	}
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
