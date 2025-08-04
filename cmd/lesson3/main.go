package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World!"
	printMe(printValue)

	var numerator int = 5
	var denominator int = 2

	fmt.Println(single_return_division_result(numerator, denominator))

	result1, remainder1, error1 := multiple_return_division_result(numerator, denominator)

	if error1 != nil {
		fmt.Printf("%s\n", error1.Error())
	} else {
		fmt.Printf("The result of the integer division is %v with %v remainder\n", result1, remainder1)
	}

	result2, remainder2, error2 := multiple_return_division_result(numerator, 0)

	// normal switch statement
	// break statement need not be explicitly mentioned
	switch {
	case error2 != nil:
		fmt.Printf("%s\n", error2.Error())
	default:
		fmt.Printf("The result of the integer division is %v with %v remainder\n", result2, remainder2)
	}

	// conditional switch statement
	switch remainder1 {
	case remainder1:
		fmt.Println("Both numbers are the exact same.")
	default:
		fmt.Println("Divison was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func single_return_division_result(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

func multiple_return_division_result(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
