package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arrN []int
	var c, numSwaps int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputN := scanner.Text()
	n, _ := strconv.Atoi(inputN)
	scanner.Scan()
	inputArr := scanner.Text()
	arr := strings.Split(inputArr, " ")
	for i := range arr {
		num, _ := strconv.Atoi(arr[i])
		arrN = append(arrN, num)
	}
	numSwaps = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arrN[j] > arrN[j+1] {
				c = arrN[j+1]
				arrN[j+1] = arrN[j]
				arrN[j] = c
				numSwaps++
			}
		}
	}
	fmt.Println("Array is sorted in", numSwaps, "swaps.")
	fmt.Println("First Element:", arrN[0])
	fmt.Println("Last Element:", arrN[len(arrN)-1])
}
