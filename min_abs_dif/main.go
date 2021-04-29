package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumAbsoluteDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func minimumAbsoluteDifference(arr []int32) (m int32) {
	// Write your code here
	// return greedyMinAbs(arr)

	return fasterMinAbs(arr)
}

func fasterMinAbs(arr []int32) (diff int32) {
	diff = 100_000_000
	// absArr(arr)
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})
	var a, b int32

	for i := 0; i < len(arr)-1; i++ {
		a = Abs(arr[i] - arr[i+1])
		b = Abs(arr[i+1] - arr[i])

		if a <= b && a < diff {
			diff = a
		} else if b < diff {
			diff = b
		}
	}

	fmt.Print(diff)
	return
}

func greedyMinAbs(arr []int32) (m int32) {
	m = 100_000_000
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if i == j {
				continue
			}
			a := Abs(arr[i] - arr[j])
			b := Abs(arr[j] - arr[i])

			if a <= b && a < m {
				m = a
			} else if b < a && b < m {
				m = b
			}
		}
	}
	fmt.Print(m)
	return
}

func Abs(i int32) int32 {
	if i < 0 {
		return i * -1
	}
	return i
}

// O(n)
func absArr(arr []int32) {
	for i := 0; i < len(arr); i++ {
		arr[i] = Abs(arr[i])
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("/home/ddowney/Workspace/github.com/hackerrank/min_abs_diff/output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := minimumAbsoluteDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
