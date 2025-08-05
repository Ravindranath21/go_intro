package main

import (
	"fmt"
)

func main() {
	var intArr [30]int32

	intArr[1] = 1
	intArr[9] = 9
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:10]) // prints elements from index [1 -> 9] both inclusive

	// initialisation of array
	var arr1 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(arr1)

	// ... syntax
	arr2 := [...]int32{1, 2, 3}
	fmt.Println(arr2)

	// slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// Output:
	// The length is 3 with capacity 3
	// The length is 4 with capacity 6

	var intSlice1 []int32 = []int32{1, 2, 3}

	// append multiple values with spread operator
	// append(originalSlice, elements...)
	intSlice = append(intSlice, intSlice1...)
	fmt.Println(intSlice)

	// if capacity is not specified then its length of the array
	// if more capacity is needed then capacity is usually doubled
	//                                       |-> capacity of slice initialised before hand for performance optimization
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// • A slice is backed by a contiguous array in RAM.
	// • When Go doubles the capacity, it tries to find a contiguous memory block of that size.
	// • If there’s not enough contiguous memory available, Go’s memory allocator:
	//  * Attempts to find a large-enough block elsewhere in the heap.
	//  * If that still fails → your program will panic with runtime error: out of memory.

	//            |-> key  |-> value type
	var myMap map[string]uint16 = make(map[string]uint16)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Steve": 24, "Hami": 25}
	fmt.Println(myMap2)

	// map returns two variables -> true if key present in map, false otherwise
	var age, ok = myMap2["Jason"]

	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Invalid name")
	}

	// Delete is done by reference so no return value
	delete(myMap2, "Hami")

	// Loops
	for name, age := range myMap2 {
		fmt.Printf("Name %v, Age %v\n", name, age)
	}

	for index, value := range intArr {
		fmt.Printf("\nIndex %v, Value %v", index, value)
	}

	i := 0
	for i < 10 {
		i += 1
	}

	i = 0
	for {
		if i > 10 {
			break
		}

		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range make([]int, 10) {
		fmt.Println(i)
	}
}
