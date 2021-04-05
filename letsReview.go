package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var arr []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	scanner.Scan()
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < n; i++ {
		in := scanner.Text()
		scanner.Scan()
		arr = append(arr, in)
	}

	for i := range arr {
		s1 := ""
		s2 := ""
		s := arr[i]
		for j := range s {
			if j%2 == 0 {
				s1 = s1 + string(s[j])
			} else {
				s2 = s2 + string(s[j])
			}
		}
		fmt.Println(s1, s2)

	}
}
