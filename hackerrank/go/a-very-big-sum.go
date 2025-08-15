// In this challenge, you need to calculate and print the sum of elements in an array, considering that some integers may be very large.

/* /*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

package main

func aVeryBigSum(ar []int64) int64 {
	var sumarr int64 = 0
	for i := 0; i < len(ar); i++ {
		sumarr += ar[i]
	}

	return sumarr

}
