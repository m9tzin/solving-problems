// Given an array of integers, calculate the ratios of its elements that are positive, negative , and zero . Print the decimal value of each fraction on a new line with 6 places after the decimal.

// Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to  10^-4 are acceptable.

package main

import (
	"fmt"
)

func plusMinus(arr []int32) {
	var positiveNums, negativeNums, zeroNums int32 = 0, 0, 0
	for _, num := range arr {
		if num > 0 {
			positiveNums++
		} else if num < 0 {
			negativeNums++
		} else if num == 0 {
			zeroNums++
		}
	}
	total := positiveNums + negativeNums + zeroNums
	fmt.Printf("%.6f\n", float64(positiveNums)/float64(total))
	fmt.Printf("%.6f\n", float64(negativeNums)/float64(total))
	fmt.Printf("%.6f\n", float64(zeroNums)/float64(total))
}
