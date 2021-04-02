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

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	res := ""
	if N%2 != 0 {
		res = res + "Weird"
	} else if N%2 == 0 {
		if (N >= 6) && (N <= 20) {
			res = res + "Weird"
		} else {
			res = res + "Not Weird"
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
