package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int64(nTemp)
	nBinary := strconv.FormatInt(n, 2) + "0"
	res := 0
	c := 0
	for i := range nBinary {
		if nBinary[i] == '1' {
			c++
		} else {
			if c > res {
				res = c
				c = 0
			} else {
				c = 0
			}
		}
	}
	fmt.Println(res)
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
