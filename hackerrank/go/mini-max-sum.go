package main

import "fmt"

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

/* arr = [1, 3, 5, 7, 9] */
/*  min = 1 + 3 + 5 + 7 = 16
max = 3 + 5 + 7 + 9 = 24 */

/* arr2 = [-1, 3, -5, 7, -9] */
/*  min = -5 + -7 + -9 = -21
    max = -1 + 3 + -5 + 7 = 4 */

func miniMaxSum(arr []int32) {
	var min, max int32 = arr[0], arr[0]
	var sum int64 = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
		sum += int64(arr[i])
	}

	fmt.Printf("%d %d\n", sum-int64(max), sum-int64(min))
}
