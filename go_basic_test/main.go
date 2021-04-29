package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

/*
 * Complete the 'Server' function below and missing types and global variables.
 *
 * The function is void.
 */

type in struct {
	value    int32
	oddChan  chan int32
	evenChan chan int32
}

var serverChan = make(chan in, 50)

func Server() {
	wg := &sync.WaitGroup{}
	for val := range serverChan {
		wg.Add(1)
		if val.value%2 == 0 {
			val.evenChan <- val.value
			wg.Done()
		} else {
			val.oddChan <- val.value
			wg.Done()
		}
		wg.Wait()
	}
}

func main() {
	stdout, err := os.Create("/home/ddowney/Workspace/github.com/hackerrank/toys/stdout.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	oddChan := make(chan int32)
	evenChan := make(chan int32)
	for idx := 0; idx < len(arr); idx++ {
		i := idx
		serverChan <- in{arr[i], oddChan, evenChan}
	}
	odds, evens := []int32{}, []int32{}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))
	go func() {
		for newOdd := range oddChan {
			odds = append(odds, newOdd)
			wg.Done()
		}
	}()
	go func() {
		for newEven := range evenChan {
			evens = append(evens, newEven)
			wg.Done()
		}
	}()
	go Server()
	wg.Wait()

	for _, resultItem := range odds {
		fmt.Fprintf(writer, "%d", resultItem)
		fmt.Fprintf(writer, "\n")
	}

	for i, resultItem := range evens {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(evens)-1 {
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
