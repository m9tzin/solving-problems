// Array Partition
// Given an integer array nums of 2n integers, group these integers into n pairs (a1, b1), (a2, b2), ..., (an, bn) such that the sum of min(ai, bi)
// for all i is maximized. Return the maximized sum.
package main

import (
	"sort"
)

// []int -> slice de inteiros (array de ints)
// func FUNCNAME(nums TYPEIN) TYPEOUT{

func arrayPairSum(nums []int) int {
	sum := 0
	// Etapa 1: Sort
	sort.Ints(nums)

	// Etapa 2: Somar elementos em posições pares (0, 2, 4, ...)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	// Etapa 3: Retornar a soma
	return sum
}
