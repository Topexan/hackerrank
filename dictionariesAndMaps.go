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
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	mp := make(map[string]string)
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	input, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	//scanner.Scan()
	for i := 0; int64(i) < input; i++ {
		//in := scanner.Text()
		//scanner.Scan()
		in := readLine(reader)
		arr := strings.Split(in, " ")
		mp[arr[0]] = arr[1]
	}
	for i := 0; int64(i) < input; i++ {
		//in := scanner.Text()
		//scanner.Scan()
		in := readLine(reader)
		_, ok := mp[in]
		if ok {
			fmt.Println(in + "=" + mp[in])
		} else {
			fmt.Println("Not found")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
