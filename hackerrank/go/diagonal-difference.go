// Given a square matrix, calculate the absolute difference between the sums of its diagonals.
/*

1 2 3
4 5 6
9 8 9

Diferença da soma das diagonais principais e secundárias: | 15 - 17 | = 2 (output)

*/
package main

import (
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	var primaryDiagSum int32 = 0
	var secondaryDiagSum int32 = 0
	for i := 0; i < len(arr); i++ {
		primaryDiagSum += arr[i][i]              // Primary Diag
		secondaryDiagSum += arr[i][len(arr)-1-i] // Secondary Diag
	}

	return int32(math.Abs(float64(primaryDiagSum - secondaryDiagSum)))
}
