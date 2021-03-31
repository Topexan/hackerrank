package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var secondInteger int
	var secondDouble float64
	var secondString string // Declare second integer, double, and String variables.
	scanner.Scan()          // Read and save an integer, double, and String to your variables.
	input1 := scanner.Text()
	secondInteger, _ = strconv.Atoi(input1)
	scanner.Scan()
	input2 := scanner.Text()
	secondDouble, _ = strconv.ParseFloat(input2, 64)
	//secondDouble, _ = strconv.Atoi(input2)

	scanner.Scan()
	secondString = scanner.Text()

	fmt.Println(i + uint64(secondInteger)) // Print the sum of both integer variables on a new line.

	fmt.Printf("%.1f\n", d+secondDouble) // Print the sum of the double variables on a new line.

	fmt.Println(s + secondString) // Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

}
