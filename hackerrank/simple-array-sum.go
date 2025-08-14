//go:build sum
// +build sum

/* Given an array of integers, find the sum of its elements. */

package main

import (
	"fmt"
)

func simpleArraySum(arr []int32) int32 {
	var sum int32 = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum) // Print the sum
	return sum
}

// Main func!
func main() {
	arr := []int32{1, 2, 3, 4, 5}
	simpleArraySum(arr)
}
