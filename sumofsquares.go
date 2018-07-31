package main

import (
	"fmt"
	"os"
)

/*
*Description:
*We want to calculate a sum of squares of some integers, excepting negatives: Sum(Ii^2 if not less than 0)
*The first line of the input will be an integer N (1 <= N <= 100)
*Each of the following N test cases consists of one line containing an integer X (0 < X <= 100), followed by X integers (Yn, -100 <= Yn <= 100) space-separated on the next line
*For each test case, calculate the sum of squares of the integers excepting negatives, and print the calculated sum to the output. No blank line between test cases
*(Take input from standard input, and output to standard output)
*Rules
*Write it in Go Programming Language
*Your source code must be a single file (package main)
*Do not use for statement
*You may only use standard library packages
 */

func sumInputs(depth int) int {
	// end recursion if end
	if depth <= 0 {
		return 0
	}

	// grab input
	var input int
	fmt.Scanf("%d", &input)

	// except negatives and check range
	if input < -100 || input > 100 {
		fmt.Println("Value to sum outside bounds.  Must be Between -100 and 100 inclusive.")
		os.Exit(1)
	}

	if input < 0 {
		input = 0
	}

	// recurse and sum square
	return sumInputs(depth-1) + (input * input)
}

func sumHandler(depth int) {
	// end recursion if end
	if depth <= 0 {
		return
	}

	// read length of number string
	var numberDepth, finalSum int
	fmt.Scanf("%d", &numberDepth)

	// throw away newline and extraneous input
	fmt.Scanln()

	// validate number
	if numberDepth > 0 && numberDepth <= 100 {
		// call recursive summing function
		finalSum = sumInputs(numberDepth)
		fmt.Scanln()
	} else {
		fmt.Println("Invalid number of integers.  Must be between 1 and 100 inclusive.")
		os.Exit(1)
	}

	// Output sum of squares
	fmt.Printf("%d%c", finalSum, '\n')

	// recurse
	sumHandler(depth - 1)
}

func main() {
	// Read in number of inputs
	var depth int
	fmt.Scanf("%d", &depth)

	// throw away newline and extraneous input
	fmt.Scanln()

	// validate number
	if depth <= 100 && depth >= 1 {
		//Call recursive sum handler, which will sum inputs
		sumHandler(depth)
	} else {
		fmt.Println("Number of entries is an invalid number. Must be between 1 and 100 inclusive.")
		return
	}
}
