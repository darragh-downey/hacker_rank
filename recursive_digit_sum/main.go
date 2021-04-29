package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
)

/*
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

type Results struct {
	Data  chan int64
	Total *big.Int
}

func NewResults() *Results {
	return &Results{
		Data:  make(chan int64),
		Total: big.NewInt(int64(0)),
	}
}

func (r *Results) Close() error {
	close(r.Data)
	return nil
}

func (r *Results) SumTotal() {
	for {
		select {
		case data := <-r.Data:
			r.Total.Add(r.Total, big.NewInt(int64(data)))
		}
	}
}

func superDigit(n string, k int32) int32 {
	// Write your code here
	startStr := ""
	for i := int32(0); i < k; i++ {
		startStr += n
	}
	var wg sync.WaitGroup
	wg.Add(1)
	results := NewResults()
	defer results.Close()

	go helper(&wg, startStr, results)
	go results.SumTotal()
	// v, _ := strconv.ParseInt(res, 10, 32)
	// fmt.Println(v)

	// return int32(v)
	return int32(results.Total.Int64())
}

func superD(n string) int32 {
	if len(n) == 1 {
		return int32(n[0] - '0')
	}
	sum := int64(0)
	nStr := make([]int32, 100)

	for _, i := range n {
		nStr = append(nStr, int32(i-'0'))
	}
	return superD(strconv.FormatInt(int64(sum), 10))
}

func helper(wg *sync.WaitGroup, n string, res *Results) {
	defer wg.Done()

	if len(n) == 1 {
		return
	}
	// sum := big.NewInt(0)

	for i := 0; i < len(n); i += 2 {
		wg.Add(1)
		// res <- sum.Add(sum, big.NewInt(int64(i-'0')))
		res.Data <- int64(n[i] - '0')
		if i == len(n)-1 {
			go helper(wg, string(n[i]), res)
		} else {
			go helper(wg, n[i:i+1], res)
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("/home/ddowney/Workspace/github.com/hackerrank/recursive_digit_sum/output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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
