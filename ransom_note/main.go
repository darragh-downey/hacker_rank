package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkMagazine(magazine []string, note []string) {
	mag := make(map[string]int)
	notes := make(map[string]int)

	for _, i := range magazine {
		if _, ok := mag[i]; !ok {
			mag[i] = 1
		} else {
			mag[i] += 1
		}
	}

	for _, i := range note {
		if _, ok := notes[i]; !ok {
			notes[i] = 1
		} else {
			notes[i] += 1
		}
	}
	// fmt.Printf("Mag count: %v\nNote word count: %v\n", mag, notes)

	for k, v := range notes {
		if _, ok := mag[k]; !ok || v > mag[k] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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
