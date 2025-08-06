package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed) // 114, uint8

	for i, v := range myString {
		fmt.Println(i, v)
	}

	// Output:
	// 0 114
	// 1 233
	// 3 115
	// 4 117
	// 5 109
	// 6 233

	// Strings use utf8 encoding

	// To overcome this problem, string can be casted as runes, which is a
	// alias for int32.

	var tempStr = []rune("Résumé")
	var ind = tempStr[1]

	fmt.Printf("%v, %T\n", ind, ind) // 82, int32

	var myRune = 'a' // runes can be intialised with single quotes
	fmt.Println(myRune)

	var strSlice = []string{"a", "b", "c"}
	var catStr = ""

	for i := range strSlice {
		catStr += strSlice[i]
	}

	// Strings are immutable in GO i.e;
	// catSlice[0] = 'b', will throw error
	// Hence in the above every iteration a new string is created.

	// This is very inefficient to tackle this problem
	// from "strings" we use strings.Builder
	// this creates an array internally and every time a addition is done
	// it appends to the string and string is created when .String() is called

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var str = strBuilder.String()

	fmt.Println(str)
}
