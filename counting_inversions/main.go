package main

import (
	"fmt"
	"math"
	"os"
)

// Complete the countInversions function below.
func countInversions(arr []int32, expected int64) (count int64) {
	count = 0
	fmt.Printf("Unsorted:\n%v\n", arr)
	if isSorted(arr) {
		fmt.Printf("Already sorted! Inversions are 0\n")
		return count
	}
	aux := make([]int32, len(arr))
	copy(aux, arr)
	count = BottomUpMerge(arr)
	fmt.Printf("Sorted:\n%v\nHad %d inversions expected %d\n", arr, count, expected)
	return count
}

// https://algs4.cs.princeton.edu/22mergesort/MergeBU.java.html
// stably merge arr[lo..mid] with arr[mid+1..hi] using temp[lo..hi]
func bumerge(arr, temp []int32, lo, mid, hi int) (count int64) {
	for k := lo; k <= hi; k++ {
		temp[k] = arr[k]
	}
	i, j := lo, mid+1
	count = 0

	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = temp[j]
			j++
		} else if j > hi {
			arr[k] = temp[i]
			i++
		} else if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
			// count will percolate through the stack
			// would need to lock.RWLock() if spawning goroutines/threads/processes
			count += int64(mid + 1 - i)
		}
	}
	return count
}

func BottomUpMerge(arr []int32) (count int64) {
	n := len(arr)
	temp := make([]int32, n)
	count = 0

	for length := 1; length < n; length *= 2 {
		for lo := 0; lo < n-length; lo += length + length {
			mid := lo + length - 1
			hi := int(math.Min(float64(lo+length+length-1), float64(n-1)))
			count += bumerge(arr, temp, lo, mid, hi)
		}
	}
	return count
}

func isSorted(arr []int32) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	stdout, err := os.Create("/home/ddowney/Workspace/github.com/hackerrank/count_inversions/output.txt")
	checkError(err)

	defer stdout.Close()

	countInversions([]int32{1, 3, 5, 7}, 0)
	countInversions([]int32{3, 2, 1}, 3)
	countInversions([]int32{1, 1, 1, 2, 2}, 0)
	countInversions([]int32{2, 1, 3, 1, 2}, 4)
	countInversions([]int32{1, 5, 3, 7}, 1)
	countInversions([]int32{7, 5, 3, 1}, 6)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
