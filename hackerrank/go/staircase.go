package main

import (
	"fmt"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	var i, j, k int32 = 0, 0, 0

	for i = 0; i < n; i++ {
		for j = 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < i+1; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}
