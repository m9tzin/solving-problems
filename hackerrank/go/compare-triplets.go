/* *
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 *

 For ex:
a = [1, 2, 3]
b = [3, 2, 1]
*/

package main

func compareTriplets(a []int32, b []int32) []int32 {
	var scoreA, scoreB int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			scoreA++
		} else if a[i] < b[i] {
			scoreB++
		}
	}
	return []int32{scoreA, scoreB}
}
