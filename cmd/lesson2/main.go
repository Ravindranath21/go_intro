package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// variables, constants and datatypes
	var intNum int = 32767 // -> int,  int8, int16, int32, int64
	intNum = intNum + 1
	fmt.Println(intNum)

	// unsigned int -> uint,  uint8, uint16, uint32, uint64

	//only float32 and float64 are there, not just float
	var floatNum float64 = 32.87
	fmt.Println(floatNum)

	// mismatched data types cannot be used to perform arthematic operations

	var myStr string = "Hello world"

	fmt.Println(len(myStr))                    // gives number of bytes occupied by string
	fmt.Println(utf8.RuneCountInString(myStr)) // this gives the length of the string

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	// all variables that are created but not initialised, GO gives defaults values
	// all int , unsigned int (uint) , float and rune -> 0
	// string -> ''
	// boolean -> false

	var myVar = "text" // type is infered
	fmt.Println(myVar)

	colonSyntax := "text" // droping the var keyword for := syntax
	fmt.Println(colonSyntax)

	// const variables need to be declared when initialised and they cannot be assigned
	// a new value later in the function
}
